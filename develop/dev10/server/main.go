package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const network = "tcp"
const address = "localhost:8080"

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("Bound to %q\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			clientReader := bufio.NewReader(conn)
			inputReader := bufio.NewReader(os.Stdin)

			go func() {
				for {
					buff, err := clientReader.ReadBytes('\n')
					if err != nil {
						return
					}
					fmt.Printf("Received: %s", buff)
				}
			}()

			for {
				buff, err := inputReader.ReadBytes('\n')
				if err != nil {
					panic(err)
				}
				if _, err := conn.Write(buff); err != nil {
					panic(err)
				}
			}
		}(conn)
	}
}
