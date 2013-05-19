package main

import (
	"fmt"
	"net/http"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func main() {
	fmt.Println("==")
	fmt.Println("==", "Go Websocket Bench")
	fmt.Println("==")
	
	host := "http://127.0.0.1:4567"
	
	req, err := http.NewRequest("GET", host, nil)
	
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(req.Method)
	
	var standard WebsocketStandard = new(WebsocketStandard_13)
	
	standard.AddHeaders(req)

	//req.Header.Add("Origin", "go-websocket-bench")

	fmt.Println(req)
	
	client := &http.Client{}
	
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(resp)
}

type WebsocketStandard interface {
	AddHeaders(req *http.Request)
	VerifyResponse(resp http.Response) bool
}

type WebsocketStandard_13 struct {}

func (s WebsocketStandard_13) AddHeaders(req *http.Request) {
	req.Header.Add("Upgrade", "websocket")
	req.Header.Add("Connection", "Upgrade")
	req.Header.Add("Sec-WebSocket-Version", "13")
	
	buf := make([]byte, 16) 
	_, err := io.ReadFull(rand.Reader, buf) 
	
	if err != nil {
		panic(fmt.Errorf("ERROR", err))
	}
	
	req.Header.Add("Sec-WebSocket-Key", base64.StdEncoding.EncodeToString(buf))
}

func (s WebsocketStandard_13) VerifyResponse(resp http.Response) (b bool) {
	return true
}

