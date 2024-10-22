package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Message struct {
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

var messages []Message
var mu sync.Mutex

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	messages = append(messages, message)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {
	receiver := r.URL.Query().Get("receiver")

	var userMessages []Message
	for _, msg := range messages {
		if msg.Receiver == receiver {
			userMessages = append(userMessages, msg)
		}
	}
	if len(userMessages) == 0 {
		http.Error(w, "No messages found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(userMessages)
}

func main() {
	http.HandleFunc("/chat/send", sendMessageHandler)
	http.HandleFunc("/chat/get", getMessagesHandler)
	log.Println("Chat Service is running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
