package main

import "log"

func main() {
	chatGPT := newChatGPT()
	message, err := chatGPT.GetHistory("Torre Eiffel")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(message)
}
