package main

import (
	"bufio"
	"encoding/hex"
	"io"
	"log"
	"net"
	"time"
)

func ConnectRemoteHost(host string) {
	NullCli := []byte{
		255,255,255,255,
		255,255,255,255,
		255,255,255,255,
		255,255,255,255,
	}

	log.Printf("Connect to %s",host)
	conn, err := net.Dial("tcp", host) // 0.0.0.0:8080
	Err(err)
	defer conn.Close()

	conn.Write(NullCli)
	b := make([]byte, 16)
	conn.Read(b)
	log.Printf("%s",hex.EncodeToString(b))
	
	for {
		id := "52fdfc072182654f163f5f0f9a621d72"
		var buff []byte
		x,_ := hex.DecodeString("")
		copy(buff, x)
		targetid,_ := hex.DecodeString(id)
		buff = append(buff, targetid...)
		buff = append(buff, '\x00')
		log.Printf("send to %s", id)
		conn.Write(buff)
		time.Sleep(3000 * time.Millisecond)
	}
}

func ConnectRemoteForDebugging(host string) {
	log.Println("Debugging Client")
	NullCli := []byte{
		255,255,255,255,
		255,255,255,255,
		255,255,255,255,
		255,255,255,255,
	}

	conn, err := net.Dial("tcp", host)
	Err(err)
	defer conn.Close()

	log.Printf("Connected to %s", host)
	conn.Write(NullCli)
	b := make([]byte, 16)
	conn.Read(b)
	log.Printf("Id: %s",hex.EncodeToString(b))
	
	for {
		buff, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			if err != io.EOF {
				Err(err)
			}
		}
		var rhuid [16]byte
		copy(rhuid[:],buff)
		log.Printf("Data from %s",hex.EncodeToString(rhuid[:]))
	}
}