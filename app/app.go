package app

import (
	"api-beer/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	router = gin.Default()
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("bad .env")
	}
}

func StartApp() {

	dbdriver := os.Getenv("DBDRIVER")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")

	models.Server.Initialize(dbdriver, username, password, port, host, database)
	fmt.Println("DATABASE STARTED")

	routes()

	router.Run(":8080")
}

