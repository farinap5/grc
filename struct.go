package main

type Message struct {
	IDSndr  string
	IDRcvr  string
	Message string
}

/*
type client struct {
	Channel chan []byte
	nick []byte
}
var a = make(map[[16]byte]client)

VVVVVVVVVVVVVV\xFFMMMMMMMMMMMMMMMMMM\x00
| ID cliente    | Nick cliente  | Mensagem para o cliente
*/