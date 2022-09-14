package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему серверу, программа должна завершаться через timeout.
*/

const network = "tcp"

func SocketClient(address string, timeout time.Duration) error {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()

	errors := make(chan error)
	gracefulShutdown := make(chan os.Signal)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	serverReader := bufio.NewReader(conn)
	inputReader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			buff, err := serverReader.ReadBytes('\n')
			if err != nil {
				errors <- err
			}
			fmt.Printf("Received: %s", buff)
		}
	}()

	go func() {
		for {
			buff, err := inputReader.ReadBytes('\n')
			if err != nil {
				errors <- err
			}
			if _, err := conn.Write(buff); err != nil {
				errors <- err
			}
		}
	}()

	select {
	case err := <-errors:
		return err
	case <-gracefulShutdown:
		fmt.Println("Exiting...")
		return nil
	}
}

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout. Default: 10 seconds")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatalln("Usage: go-telnet [--timeout=] <host> <port>")
	}
	address := net.JoinHostPort(args[0], args[1])
	if err := SocketClient(address, *timeout); err != nil {
		panic(err)
	}
}
