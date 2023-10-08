package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"io"
	"log"
	"net"
	//"time"
)

// var Clients = make(map[[16]string]chan string)
var Clients = make(map[[16]byte]chan []byte)

func (m Message)Send(message string) {
	m.Message = message
	var buff []byte
	for i := 0; i < 16; i++ {
		buff = append(buff, m.IDSndr[i])
	}
	buff = append(buff, []byte(m.Message)...)
	buff = append(buff, '\x00')
	Clients[m.IDRcvr] <- buff
  }
  

func NewMessage(from, to [16]byte) Message {
  m := Message{
    IDSndr: from,
    IDRcvr: to,
  }
  return m
}


func HandleClient(uid [16]byte, conn net.Conn, v int) {
	go func() {
	  for {
		/*
			This loop id for my client to send messages
			The server receives data from socker, will deserialize it and pass to the client
			receive data from socket (my client)
		*/
	   	buff, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			if err == io.EOF {
				println(err.Error())
				break
			} else {
				println(err.Error())
				break
			}
		}

		if v > 2 {
			log.Printf("Got buffer from %s %s",conn.RemoteAddr().String(),hex.EncodeToString(uid[:]))
		}
		var toId [16]byte
		copy(toId[:],buff)
		if v > 3 {
			log.Printf("+ send to %s",hex.EncodeToString(toId[:]))
		}
		m := NewMessage(uid,toId)
		m.Send("message") 
	  }
	}()
	go func() {
		for {
			msg := <-Clients[uid]
				if v > 3 {
						log.Printf("+ received %s ", hex.EncodeToString(uid[:]))
					}
					_, err := conn.Write(msg)
					println("aaaaa")
					Err(err)
			}
			println("xxx")
	}()
	//log.Printf("Remote host %s %s id done", conn.RemoteAddr().String(), hex.EncodeToString(uid[:]))
}


/*
  If the client has a ID assigned he must start the conversation
*/
func ConnectClient(conn net.Conn) ([16]byte,error) {
	var uid [16]byte
	NullCli := []byte{
		255,255,255,255,
		255,255,255,255,
		255,255,255,255,
		255,255,255,255,
	}
	
	/*
		The id must be passed ending with null byte
		like: byte(blablabla)\x00
		id: 16 bytes
	*/
	//data, err := bufio.NewReader(conn).ReadBytes(0x00)
	// Byffer to store ID of the reote host
	buff := make([]byte, 16)
	// Read from socket
	_, err := conn.Read(buff)
	if err != nil {
		return uid,err
	}
	// if data from socket is not a valid ID, create one
	// invalid id == \xff * 16
	if bytes.Equal(NullCli,buff) {
		uid = GenRandomData()
		log.Printf("Remote host %s assigned %s",conn.RemoteAddr().String(),hex.EncodeToString(uid[:]))
	} else {
		// convert buffer (id) from ID bytes to string hex equivalent 
		copy(uid[:],buff)
		log.Printf("Remote host %s known as %s",conn.RemoteAddr().String(),uid)
	}

	clientChann := make(chan []byte)
	Clients[uid] = clientChann
	return uid,nil
}

func StartServer(host string, verbosity int) {
	bind, err := net.Listen("tcp", host)
	Err(err)

	if verbosity > 0 {
		log.Printf("Listening on %s", host)
	}

	defer bind.Close()
	for {
		conn, err := bind.Accept()
		
		Err(err)
	  	if verbosity > 0 {
			log.Printf("Remote host %s connected", conn.RemoteAddr().String())
	  	}

	  	uid, err := ConnectClient(conn)
	  	if err != nil  {
			if err != io.EOF {
				log.Printf("Remote host %s closed", conn.RemoteAddr().String())
				Err(err)
			} else {
				log.Printf("Remote host %s EOF", conn.RemoteAddr().String())
			}
	  	}
		// send ID to client
	  	conn.Write(uid[:])

		// fvck
		go HandleClient(uid,conn,verbosity)
	}
}
