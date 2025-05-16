package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type OtusTelnetClient struct {
	address       string
	timeout       time.Duration
	in            io.ReadCloser
	out           io.Writer
	conn          net.Conn
	ctx           context.Context
	ctxCancelFunc context.CancelFunc
}

func NewOtusTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) *OtusTelnetClient {
	ctx, cancel := context.WithCancel(context.Background())
	return &OtusTelnetClient{
		address:       address,
		timeout:       timeout,
		in:            in,
		out:           out,
		ctx:           ctx,
		ctxCancelFunc: cancel,
	}
}

func (cl *OtusTelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", cl.address, cl.timeout)
	if err != nil {
		return fmt.Errorf("соединение: %w", err)
	}
	cl.conn = conn
	return nil
}

func (cl *OtusTelnetClient) Close() error {
	err := cl.conn.Close()
	if err != nil && !errors.Is(err, net.ErrClosed) {
		return fmt.Errorf("ошибка при закрытии соединения: %w", err)
	}
	return nil
}

func (cl *OtusTelnetClient) Send() error {
	scanner := bufio.NewScanner(cl.in)
	for scanner.Scan() {
		_, err := fmt.Fprintln(cl.conn, scanner.Text())
		if err != nil {
			return fmt.Errorf("отправка данных в соединение: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("чтения ввода: %w", err)
	}

	return nil
}

func (cl *OtusTelnetClient) Receive() error {
	scanner := bufio.NewScanner(cl.conn)
	for scanner.Scan() {
		_, err := fmt.Fprintln(cl.out, scanner.Text())
		if err != nil {
			return fmt.Errorf("не удалось получить сообщение: %w", err)

		}
	}
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return NewOtusTelnetClient(address, timeout, in, out)
}
