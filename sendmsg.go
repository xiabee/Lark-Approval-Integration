package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SendMsg(msg string) {
	secret := os.Getenv("WEBHOOK_KEY")
	webhookURL := "https://open.feishu.cn/open-apis/bot/v2/hook/" + secret

	// POST to Lark Bot
	message := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]interface{}{
			"text": msg,
		},
	}

	payload, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Failed to serialize message:", err)
		return
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Failed to send message:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Message sent successfully")
	} else {
		fmt.Println("Failed to send message. Status:", resp.StatusCode)
	}
}
