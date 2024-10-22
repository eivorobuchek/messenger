package main

import (
	"auth_service/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Api struct {
	router *gin.Engine
}

func main() {

	var api Api
	api.registerRoutes()

	err := api.router.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Auth Service is running on port 8081")
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

	auth := base.Group("/auth")
	auth.POST("/register", handlers.RegisterHandler)
	auth.POST("/login", handlers.LoginHandler)

	api.router = router
}
