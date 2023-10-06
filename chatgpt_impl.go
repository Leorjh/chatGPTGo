package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type chatImpl struct {
	ApiKey string
	URL    string
}

func newChatGPT() IADoer {
	return chatImpl{
		ApiKey: "API-KEY",
		URL:    "https://api.openai.com/v1/chat/completions",
	}
}

type RequestChatGPT struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

func (c chatImpl) GetHistory(text string) (string, error) {
	input := RequestChatGPT{
		Model:  "gpt-3.5-turbo",
		Prompt: text,
	}

	byteResult, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(byteResult))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}
