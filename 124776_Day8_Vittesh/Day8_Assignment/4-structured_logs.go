package main

import (
	"encoding/json"
	"log"
)

type LogMessage struct {
	Name   string `json:"Name"`
	Address string `json:"Address"`
}

func main() {

	info := LogMessage{
		Name:   "Vittesh",
		Address: "Bengaluru",
	}

	data, _ := json.Marshal(info)

	log.Println(string(data))
}