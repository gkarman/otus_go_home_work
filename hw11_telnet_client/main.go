package main

import (
	"fmt"
	"log"
	"net"
	"os"
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		writeRoutine(client)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		readRoutine(client)
	}()
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

func writeRoutine(client TelnetClient) {
	err := client.Send()
	if err != nil {
		log.Fatalf("отправка не удалась %v", err)
	}
}

func readRoutine(client TelnetClient) {
	err := client.Receive()
	if err != nil {
		log.Fatalf("получение не удалась %v", err)
	}
}

func closeClient(client TelnetClient) {
	if err := client.Close(); err != nil {
		log.Printf("ошибка при закрытии соединения: %v", err)
	}
}
