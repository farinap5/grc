package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"time"
)

func ConnectRemoteHost(host string) {
	log.Printf("Connect to %s",host)
	conn, err := net.Dial("tcp", host) // 0.0.0.0:8080
	Err(err)
	defer conn.Close()

	conn.Write([]byte("elf"+"\x00"))

	data, err := bufio.NewReader(conn).ReadBytes(0x00)
	if err != nil {
		log.Printf("%s", err.Error())
	}
	log.Printf("%s", string(data))
	
	for {
		id := "gnome" // target id
		var buff []byte

		buff = append(buff,	[]byte(id)...)
		buff = append(buff, '\xff')
		buff = append(buff, []byte("probe")...)
		buff = append(buff, '\x00')
		log.Printf("send to %s", id)
		conn.Write(buff)
		time.Sleep(3000 * time.Millisecond)
	}
}

func inv(s string) (x string) {
    for _, v := range s {
        x = string(v) + x
    }
    return
}

func ConnectRemoteForDebugging(host string) {
	log.Println("Debugging Client")

	username := "bot"
	conn, err := net.Dial("tcp", host)
	Err(err)
	defer conn.Close()

	log.Printf("Connected to %s", host)
	conn.Write([]byte(username+"\x00")) // send local username
	data, err := bufio.NewReader(conn).ReadBytes(0x00) // receive "username_connected"
	if err != nil {
		log.Printf("%s", err.Error())
	}
	log.Printf("%s", string(data))
	
	for {
		data, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			if err != io.EOF {
				Err(err)
			}
		}
		senderName := GetUserName(data)
		message, err := GetMessage(data)
		if err != nil {
			log.Printf("%s",err.Error())
		}

		log.Printf("Message from %s : %s",string(senderName), string(message))
		if conn == nil {
			log.Printf("Connection Closed")
			break
		}

		message = message[1:]
		return_message := inv(string(message))
		
		var buff []byte
		buff = append(buff, []byte(senderName)...)
		buff = append(buff, '\xff')
		buff = append(buff, []byte(return_message)...)
		buff = append(buff, '\x00')
		conn.Write(buff)
	}
}