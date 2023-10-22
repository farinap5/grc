package main

import (
	"bufio"
	//"fmt"
	"io"
	"log"
	"net"
	//"time"
)

// var Clients = make(map[[16]string]chan string)
var Clients = make(map[string]chan []byte)

func (m Message)Send(message string) {
	m.Message = message
	var buff []byte

	buff = append(buff, []byte(m.IDSndr)...)
	//buff = append(buff, '\xff')
	buff = append(buff, []byte(m.Message)...)
	buff = append(buff, '\x00')
	//fmt.Println(buff)
	//fmt.Println([]byte(m.IDRcvr))
	Clients[m.IDRcvr] <- buff
}

func NewMessage(from, to string) Message {
  m := Message{
    IDSndr: from,
    IDRcvr: to,
  }
  return m
}

func HandleClient(uid string, conn net.Conn, v int) {
	//println(uid)
	//fmt.Println([]byte(uid))
	go func() {
	  for {
		/*
			This loop id for my client to send messages
			The server receives data from socker, will deserialize it and pass to the client
			receive data from socket (my client)

			Stream -> name\xffmessage\x00
		*/
	   	data, err := bufio.NewReader(conn).ReadBytes(0x00)
		if err != nil {
			if err == io.EOF {
				log.Printf("Remote host %s EOF receive loop",conn.RemoteAddr())
				break
			} else {
				log.Printf("Remote host %s close due %s", conn.RemoteAddr(), err.Error())
				break
			}
		}

		toName := GetUserName(data)
		message,err := GetMessage(data)
		if err != nil {
			log.Printf("Remote host %s error %s", conn.RemoteAddr(), err.Error())
			break
		}

		if v > 2 {
			log.Printf("Got buffer from %s %s", conn.RemoteAddr().String(), uid)
		}
		if (toName[0] == '\x01') {
			log.Printf("Broadcast issued ");
			for toNameString := range Clients {
				m := NewMessage(uid, toNameString)
				m.Send(string(message))
			}  

		} else {
			toNameString := string(toName)
			if v > 3 {
				log.Printf("+ send to %s", toNameString)
			}
			m := NewMessage(uid,toNameString)
			m.Send(string(message)) 
		}
		
	  }
	}()
	for {
		msg := <-Clients[uid]
		//println(msg)
		if v > 3 {
			log.Printf("+ received %s ", uid)
		}
		_, err := conn.Write(msg)
		if err != nil {
			log.Printf("Remote host %s %s",uid,err.Error())
			break
		}
	}
	println("xxx")
	//log.Printf("Remote host %s %s id done", conn.RemoteAddr().String(), hex.EncodeToString(uid[:]))
}


/*
  If the client has a ID assigned he must start the conversation
*/
func ConnectClient(conn net.Conn) (string,error) {
	Name, err := bufio.NewReader(conn).ReadBytes(0x00)
	if err != nil {
		return "", err
	}
	
	/*
		Receive the username from the socket
		Create a channel and map it to be callable from the name

		receive: elf\x00
		turn into a string erasing the null byte ending.
	*/

	NameString := string(Name[:len(Name)-1])
	clientChann := make(chan []byte)
	Clients[NameString] = clientChann
	return NameString, nil
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

	  	uid, err := ConnectClient(conn) // uid = name
	  	if err != nil  {
			if err != io.EOF {
				log.Printf("Remote host %s closed", conn.RemoteAddr().String())
				Err(err)
			} else {
				log.Printf("Remote host %s EOF 1", conn.RemoteAddr().String())
			}
	  	} else {
			if verbosity > 0 {
				log.Printf("+ %s", uid)
			}
			// send ID to client
			  conn.Write([]byte(uid+"_connected"+"\x00"))
	
			// fvck
			go HandleClient(uid,conn,verbosity)
		}
	}
}
