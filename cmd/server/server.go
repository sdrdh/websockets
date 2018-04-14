package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

type GameMap map[string]map[string]chan string

var gameMap GameMap

// Messages:
// userid: userid
// start_room:
// join_room:roomNo
// msg: msgText

func writeToConn(conn net.Conn, msgChan chan string) {
	for {
		v := <-msgChan
		fmt.Fprintln(conn, v)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Got A Connection")
	reader := bufio.NewReader(conn)
	var userId string
	var roomIdStr string
	for {
		msg, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		if err == io.EOF {
			break
		}
		if msg == "\n" {
			continue
		}
		lexemes := strings.Split(msg, ":")
		fmt.Println(len(lexemes))
		fmt.Println("Case Switch:", lexemes[0])
		switch lexemes[0] {
		case "userid":
			userId = lexemes[1]
			continue
		case "start_room":
			roomIdStr = lexemes[1]
			fmt.Println("Creating the Room:", roomIdStr)
			userChan := make(chan string)
			go writeToConn(conn, userChan)
			roomMap := make(map[string]chan string)
			roomMap[userId] = userChan
			gameMap[roomIdStr] = roomMap
			continue
		case "join_room":
			roomIdStr = lexemes[1]
			fmt.Printf("User %s joining the Room %s", userId, roomIdStr)
			userChan := make(chan string)
			go writeToConn(conn, userChan)
			roomMap, exists := gameMap[roomIdStr]
			if !exists {
				fmt.Println("Room doesnt exist")
				continue
			}
			roomMap[userId] = userChan
			continue
		case "msg":
			fmt.Println("Hi I usually come here all the time")
			roomMap := gameMap[roomIdStr]
			for k, v := range roomMap {
				if k != userId {
					v <- lexemes[1]
				}
			}
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	gameMap = make(GameMap)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}
