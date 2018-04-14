package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func readFromStdIn(reader *bufio.Reader, ioChan chan string) {
	for {
		text, _ := reader.ReadString('\n')
		ioChan <- text
	}
}

func readFromConn(reader *bufio.Reader, connChan chan string) {
	for {
		status, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if status == "\n" {
			continue
		}
		connChan <- status
	}
}

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	connreader := bufio.NewReader(conn)
	ioreader := bufio.NewReader(os.Stdin)
	ioChan := make(chan string)
	connChan := make(chan string)
	go readFromStdIn(ioreader, ioChan)
	go readFromConn(connreader, connChan)
	for {
		select {
		case v := <-ioChan:
			fmt.Fprint(conn, fmt.Sprintf("%s\n", v))
		case v := <-connChan:
			fmt.Println("Message from conn:", v)
		}
	}
}
