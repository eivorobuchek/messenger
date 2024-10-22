package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func RegisterHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func LoginHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Success login"})
}
