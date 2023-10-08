package main

import (
	"log"
	"math/rand"
	"os"
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