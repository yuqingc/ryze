package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// IsValidToken checks if a token is valid
func IsValidToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my-secret"), nil
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false
}

// HandleVarifyToken is user for /api/veryfy_token
func HandleVarifyToken(c *gin.Context) {
	userToken := c.GetHeader("Authorization")
	fmt.Println("usertoken, ", userToken)
	trimmedTokenString := strings.TrimPrefix(userToken, `Bearer `)
	trimmedTokenString = strings.TrimSpace(trimmedTokenString)
	if IsValidToken(trimmedTokenString) {
		c.String(http.StatusOK, "access_token is valid")
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
