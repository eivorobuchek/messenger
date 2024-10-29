package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserProfile struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}

var profiles []UserProfile

func UpdateProfileHandler(c *gin.Context) {
	var profile UserProfile

	if err := json.NewDecoder(c.Request.Body).Decode(&profile); err != nil {
		http.Error(c.Writer, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, p := range profiles {
		if p.Username == profile.Username {
			profiles[i] = profile
			c.Writer.WriteHeader(http.StatusOK)
			json.NewEncoder(c.Writer).Encode(profile)
			return
		}
	}
	profiles = append(profiles, profile)
	c.Writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(c.Writer).Encode(profile)
}

func SearchProfileHandler(c *gin.Context) {
	username := c.Request.URL.Query().Get("username")

	for _, profile := range profiles {
		if profile.Username == username {
			json.NewEncoder(c.Writer).Encode(profile)
			return
		}
	}
	http.Error(c.Writer, "User not found", http.StatusNotFound)
}
