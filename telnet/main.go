package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("usage: telnet host port [timeout=10s]")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("connection error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("connected to", address)

	done := make(chan struct{})

	// чтение из сокета
	go func() {
		defer close(done)
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println("read error:", err)
		}
	}()

	// ввод пользователя
	go func() {
		reader := bufio.NewReader(os.Stdin)

		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF { // Ctrl+D
				conn.Close()
				return
			}
			if err != nil {
				fmt.Println("input error:", err)
				return
			}

			_, err = conn.Write(line)
			if err != nil {
				return
			}
		}
	}()

	<-done
	fmt.Println("\nconnection closed")
}
