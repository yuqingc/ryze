package main

import (
	"fmt"

	"github.com/yuqingc/ryze/pkg/signup"
)

func main() {
	writeRes, err := signup.WriteUserToDB("admin", "pwd123456")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(writeRes)
}
