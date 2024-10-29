package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

func SendMessageHandler(c *gin.Context) {
	var message Message
	if err := json.NewDecoder(c.Request.Body).Decode(&message); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid input")
		return
	}

	mu.Lock()
	messages = append(messages, message)
	mu.Unlock()

	c.Writer.WriteHeader(http.StatusCreated)
	c.JSON(http.StatusOK, messages)
}

func GetMessagesHandler(c *gin.Context) {
	receiver := c.Request.URL.Query().Get("receiver")

	var userMessages []Message
	for _, msg := range messages {
		if msg.Receiver == receiver {
			userMessages = append(userMessages, msg)
		}
	}
	if len(userMessages) == 0 {
		c.JSON(http.StatusNotFound, "No messages found")
		return
	}
	c.JSON(http.StatusOK, userMessages)
}
