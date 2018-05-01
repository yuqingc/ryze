package auth

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
	nowSec := time.Now().Unix()
	expireSec := int64(3600 * 24)
	var dbUsername, dbPassword string
	db, err := db.OpenMySQL()
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer db.Close()
	// TODO: prevent SQL injection (done)
	passwdRow, err := db.Query("select * from passwd where username=?", cUsername)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer passwdRow.Close()
	for passwdRow.Next() {
		err = passwdRow.Scan(&dbUsername, &dbPassword)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		// TODO: prevent sql injection & check username with token?
		fmt.Println(dbUsername, dbPassword)

	}
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(cPassword))
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusUnauthorized, err.Error())
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "Ryze/general",
		"exp": nowSec + expireSec,
		"iat": nowSec,
		"usr": cUsername,
	})
	mySecret := []byte("my-secret")
	tokenString, err := token.SignedString(mySecret)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"username":     cUsername,
			"token_type":   "JWT",
			"access_token": tokenString,
			"expires_in":   expireSec,
		})
	} else {
		c.String(http.StatusBadRequest, err.Error())
	}
}
