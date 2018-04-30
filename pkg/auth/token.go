package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// IsValidToken checks if a token is valid. If it is valid, the token claims is returned
func IsValidToken(tokenString string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my-secret"), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	if tokenClaim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return tokenClaim, true
	}

	return nil, false
}

// HandleVarifyToken is user for /api/veryfy_token
func HandleVarifyToken(c *gin.Context) {
	userToken := c.GetHeader("Authorization")
	fmt.Println("usertoken, ", userToken)
	trimmedTokenString := strings.TrimPrefix(userToken, `Bearer `)
	trimmedTokenString = strings.TrimSpace(trimmedTokenString)
	if tokenClaims, ok := IsValidToken(trimmedTokenString); ok {
		fmt.Println("tokenClaims is ", tokenClaims)
		c.String(http.StatusOK, "access_token is valid")
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
