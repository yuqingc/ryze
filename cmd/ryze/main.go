package main

import (
	"fmt"

	"github.com/yuqingc/ryze/pkg/auth"
	"github.com/yuqingc/ryze/pkg/hello"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Ryze is going...")
	sayhi()

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/hello", hello.HandleHello)
		api.POST("/login", auth.HandleLogin)
		api.GET("/varify_token", auth.HandleVarifyToken)
	}

	router.Run(":8080")
}
