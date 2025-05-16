package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/spf13/pflag"
)

const DefaultTimeout = 10 * time.Second

type args struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func main() {
	args, err := parseArgs()
	if err != nil {
		log.Fatalf("ошибка формирование конфига %v", err)
	}

	address := net.JoinHostPort(args.Host, args.Port)
	client := NewTelnetClient(address, args.Timeout, os.Stdin, os.Stdout)
	defer closeClient(client)
	if err := client.Connect(); err != nil {
		log.Fatalf("подключение не удалось %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer closeClient(client)
		writeRoutine(ctx, client)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer closeClient(client)
		readRoutine(ctx, client)
	}()

	<-ctx.Done()
	_ = client.Close()

	wg.Wait()
}

func parseArgs() (args, error) {
	var timeout time.Duration
	pflag.DurationVar(&timeout, "timeout", DefaultTimeout, "таймаут подключения")
	pflag.Parse()

	a := pflag.Args()
	if len(a) < 2 {
		return args{}, fmt.Errorf("неверные параметры")
	}
	host := a[0]
	port := a[1]

	return args{host, port, timeout}, nil
}

func writeRoutine(ctx context.Context, client TelnetClient) {
	errorCh := make(chan error, 1)
	go func() {
		errorCh <- client.Send()
	}()

	select {
	case <-ctx.Done():
		return
	case err := <-errorCh:
		if err != nil {
			return
		}
	}
}

func readRoutine(ctx context.Context, client TelnetClient) {
	errorCh := make(chan error, 1)
	go func() {
		errorCh <- client.Receive()
	}()

	select {
	case <-ctx.Done():
		log.Println("остановка чтения (Ctrl+C)")
		return
	case err := <-errorCh:
		if err != nil {
			log.Printf("ошибка при получении: %v", err)
		}
		return
	}
}

func closeClient(client TelnetClient) {
	if err := client.Close(); err != nil {
		log.Printf("ошибка при закрытии соединения: %v", err)
	}
}
