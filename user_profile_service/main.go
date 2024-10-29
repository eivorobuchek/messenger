package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"user_profile_service/internal/handlers"
)

type Api struct {
	router *gin.Engine
}

func main() {

	var api Api
	api.registerRoutes()

	err := api.router.Run(":8082")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("User Profile Service is running on port 8082")
}

const BasePath = "/api"

func (api *Api) registerRoutes() {
	router := gin.Default()

	// Liveness probe: сервис работает
	router.GET("/health/liveness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "alive"})
	})

	// Readiness probe: сервис готов принимать трафик
	router.GET("/health/readiness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	base := router.Group(BasePath)

	auth := base.Group("/profile")
	auth.POST("/update", handlers.UpdateProfileHandler)
	auth.POST("/search", handlers.SearchProfileHandler)

	api.router = router
}
