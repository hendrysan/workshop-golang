package main

import (
	"log"
	"workshop1/config"
	"workshop1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitViper()

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err.Error())
		// panic(err)
	} else {
		log.Println("Configuration loaded successfully")
	}

	db, err := config.InitDB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
		// panic(err)
	} else {
		log.Println("Database connection established successfully")
	}

	router := gin.Default()
	routes.Routes(router, db)

	router.Run(":8080")
}
