package login

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yuqingc/ryze/pkg/db"
)

// HandleLogin is used for POST /api/login
func HandleLogin(c *gin.Context) {
	cUsername := c.PostForm("username")
	cPassword := c.PostForm("password")
	var dbPassword string
	db, err := db.OpenMySQL()
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()
	queryBody := fmt.Sprintf("select password from passwd where username='%s'", cUsername)
	passwdRow, err := db.Query(queryBody)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer passwdRow.Close()
	for passwdRow.Next() {
		err = passwdRow.Scan(&dbPassword)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

	}
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(cPassword))
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusUnauthorized, err.Error())
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "Ryze/general",
		"exp": 3600 * 24,
		"iat": time.Now().Unix(),
		"usr": cUsername,
	})
	mySecret := []byte("my-secret")
	tokenString, err := token.SignedString(mySecret)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"username":     cUsername,
			"access_token": tokenString,
		})
	} else {
		c.String(http.StatusBadRequest, err.Error())
	}
}
