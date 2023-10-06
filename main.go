package main

import (
    "fmt"
    "net/http"
    "bytes"
    "io"
)

func main() {
    apiKey := "YOUR-KEY"

    apiUrl := "https://api.openai.com/v1/engines/gpt-3.5-turbo-instruct/completions"

    input := `{
        "prompt": "\"História torre eifel",
        "max_tokens": 50
    }`

    req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer([]byte(input)))
    if err != nil {
        panic(err)
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    // Exibe a resposta JSON
    fmt.Println(string(responseBody))
}
