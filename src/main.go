package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var msg *Message = &Message{
			Sender:   "Vítor Mariano",
			Receiver: "Andrey Vital",
			Message:  "Já pode parar de chorar. Este é meu primeiro código em go.",
		}

		j, err := json.Marshal(msg)

		if err != nil {
			log.Fatal(err)
		}

		w.Write(j)
	})

	log.Println("Serving on localhost:3000")
	http.ListenAndServe(":3000", nil)
}
