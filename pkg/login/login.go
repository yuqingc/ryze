package login

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	username := c.PostForm("username")
	// password := c.PostForm("password")
	fmt.Println("name is ", username)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
	})
	tokenString, err := token.SignedString("secret")
	fmt.Println("xxx ", tokenString, err)
	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
	})
}
