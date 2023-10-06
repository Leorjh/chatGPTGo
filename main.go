package main

import "log"

func main() {
	ia := newChatGPT()
	message, err := ia.GetHistory("Torre Eiffel")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(message)
}
