package main

import (
	"fmt"

	"github.com/yuqingc/ryze/pkg/hello"
	// "github.com/yuqingc/ryze/pkg/signup"
	"github.com/yuqingc/ryze/pkg/login"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("main")
	fk()
	// const pswd string = "password123456"
	// var data []byte = []byte(pswd)
	// result, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("jiami", string(result))
	// }
	// h1 := []byte("$2a$10$WvY1m3O6MMET2mwLjSvQXOGYE7gYGxyaQFf53Ds1t7mEFEQwEwK.W")
	// h2 := []byte("$2a$10$di4i9HPY8p17hwy6a/N6B.r4FRSadQUuP9RzWdnoT3nB5P3muEoY2")
	// err = bcrypt.CompareHashAndPassword(h1, data)
	// if err == nil {
	// 	fmt.Println("h1, compare successfully!")
	// } else {
	// 	fmt.Println(err)
	// }

	// err = bcrypt.CompareHashAndPassword(h2, data)
	// if err == nil {
	// 	fmt.Println("h2, compare successfully!")
	// } else {
	// 	fmt.Println(err)
	// }

	// writeRes, err := signup.WriteUserToDB("admin", "pwd123123")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(writeRes)

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/hello", hello.HandleHello)
		api.POST("login", login.HandleLogin)
	}

	router.Run(":8080")
}
