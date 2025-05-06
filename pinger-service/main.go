package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	data "github.com/Robert076/logger/pinger-service/messages"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		randomMessage := data.Messages[rand.Intn(len(data.Messages))]

		message := struct {
			Message string `json:"message"`
		}{
			Message: randomMessage,
		}

		messageJSON, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error marshaling message:", err)
			return
		}

		url := "http://localhost:1234/ping"
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(messageJSON))
		if err != nil {
			fmt.Println("Error sending HTTP request:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println("Sent message:", randomMessage)
		fmt.Println("Response Status:", resp.Status)
	}
}
