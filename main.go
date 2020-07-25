package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := NewMessage(WithAddr("北京"), WithID(12))
	fmt.Println(m)
	bytes, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

func NewMessage(opt ...MessageOption) Message {
	m := Message{}
	for _, o := range opt {
		o(&m) //TODO 重点
	}
	return m
}

type Message struct {
	Id      int
	Name    string
	Address string
	Phone   int
}

type MessageOption func(msg *Message) // TODO 重点

func WithID(id int) MessageOption {
	return func(m *Message) { //TODO 重点
		m.Id = id
	}
}

func WithName(name string) MessageOption {
	return func(m *Message) { m.Name = name }
}
func WithAddr(addr string) MessageOption {
	return func(m *Message) { m.Address = addr }
}
