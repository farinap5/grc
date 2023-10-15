package main

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"strings"
)

func Err(err error) {
	if err != nil {
	  log.Printf("%s\n",err.Error())
	  os.Exit(1)
	}
}

func GenRandomData() [16]byte {
	b := make([]byte, 16)
	var xxx [16]byte
	_, err := rand.Read(b)
	if err != nil {
		Err(err)
	}
	copy(xxx[:],b)
	return xxx
}

func GetUserName(data []byte) []byte {
	ff := strings.Index(string(data), "\xff")
    if (ff < 0) {
        log.Println("Error: no \\xff terminator!")
		return nil
    }
    return data[:ff]
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