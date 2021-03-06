package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/richiMarchi/scratchpay-challenge/internal/core/services"
	"github.com/richiMarchi/scratchpay-challenge/internal/handlers"
	"github.com/richiMarchi/scratchpay-challenge/internal/repositories"
)

const DefaultCertPath = "./misc/server.crt"
const DefaultPvtKeyPath = "./misc/server.key"

func main() {
	usersRepository := repositories.NewMySqlDb(
		getEnv("DB_USER", "tech-user"), getEnv("DB_PASS", "tech-pw"), getEnv("DB_HOST", "localhost"), getEnv("DB_NAME", "usersdb"))
	usersService := services.New(usersRepository)
	usersHandler := handlers.New(usersService)

	router := gin.New()
	router.GET("/users/:userId", usersHandler.GetUser)
	router.POST("/users", usersHandler.CreateUser)
	router.GET("/users", usersHandler.ListUsers)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	router.RunTLS(":8080", getEnv("CERT_PATH", DefaultCertPath), getEnv("PVTKEY_PATH", DefaultPvtKeyPath))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
