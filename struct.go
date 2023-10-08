package main

type Message struct {
	IDSndr  [16]byte
	IDRcvr  [16]byte
	Message string
}