package chatsrc

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"strings"

	//"github.com/marcusolsson/tui-go"
)

var connx net.Conn

func InitConnection(host string, id string) (string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return err.Error()
	}

	conn.Write([]byte(id+"\x00"))
	data, err := bufio.NewReader(conn).ReadBytes(0x00)
	if err != nil {
		return err.Error()
	}
	if string(data) == id+"_connected"+"\x00" {
		User = id
		connx = conn
		go Binder()
		return string(data)
	} else {
		return "Error connecting"
	}
}

func Send(id, msg string) string {
	if connx == nil {
		return "Nil socket"
	}
	var buff []byte
	buff = append(buff, []byte(id)...)
	buff = append(buff, '\xff')
	buff = append(buff, []byte(msg)...)
	buff = append(buff, '\x00')
	connx.Write(buff)
	return msg
}

func CloseSocket() {
	connx.Close()
}

func Binder() {
	PrintMessage("system","Starting binder")
	for {
		data, err := bufio.NewReader(connx).ReadBytes(0x00)
		if err != nil {
			if err != io.EOF {
					PrintMessage("system","End Of File. Closed!")
			} else {
				PrintMessage("system",err.Error())
			}
		} else {
			senderName := GetUserName(data)
			message, err := GetMessage(data)
			if err != nil {
				PrintMessage("system",err.Error())
			} else {
				PrintMessage(string(senderName),
				string(message))
			}
		}
	}
}

func GetMessage(data []byte) ([]byte, error) {
	ff := strings.Index(string(data), "\xff")
    if (ff < 0) {
		return nil, errors.New("no \\xff terminator")
    }
	message := data[ff:]

	nulli := strings.Index(string(message), "\x00")
    if (ff < 0) {
        log.Println("Error: no null terminator!")
		return nil, errors.New("no \\xff terminator")
    }
    return message[:nulli],nil
}

func GetUserName(data []byte) []byte {
	ff := strings.Index(string(data), "\xff")
    if (ff < 0) {
        log.Println("Error: no \\xff terminator!")
		return nil
    }
    return data[:ff]
}