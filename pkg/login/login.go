package login

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	username := c.PostForm("username")
	// password := c.PostForm("password")
	fmt.Println("name is ", username)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2018, 4, 25, 20, 0, 0, 0, time.UTC).Unix(),
	})
	mySecret := []byte("my-secret")
	tokenString, err := token.SignedString(mySecret)
	fmt.Println("token ", tokenString, "err ", err)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"access_token": tokenString,
		})
	} else {
		c.String(http.StatusForbidden, err.Error())
	}
}
