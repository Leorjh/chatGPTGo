package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

func main() {
	apiKey := "API-KEY"
	apiUrl := "https://api.openai.com/v1/chat/completions"

	input := Request{
		Model:  "gpt-3.5-turbo",
		Prompt: "Torre Eiffel",
	}

	byteResult, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(byteResult))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseBody))
}
