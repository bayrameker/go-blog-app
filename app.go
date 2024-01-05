package main

import (
	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("failed to load .env file")
	}

	err = db.InitializeDB()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")

}
