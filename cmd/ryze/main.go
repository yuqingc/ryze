package main

import (
	"fmt"

	"github.com/yuqingc/ryze/pkg/auth"
	"github.com/yuqingc/ryze/pkg/hello"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("main")
	fk()

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/hello", hello.HandleHello)
		api.POST("login", auth.HandleLogin)
	}

	router.Run(":8080")
}
