package main

import (
	"log"
	"os"

	"github.com/bajra-manandhar17/personal-finance-app/cmd/http/get_user"
	"github.com/bajra-manandhar17/personal-finance-app/cmd/http/sign_up"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/apis/users/signup", sign_up.RegisterNewUserHandler)
	r.GET("/apis/users/:userId", get_user.GetUserHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
