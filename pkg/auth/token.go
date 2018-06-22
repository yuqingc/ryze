package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// isValidToken checks if a token is valid. If it is valid, the token claims is returned
func isValidToken(tokenString string) (jwt.MapClaims, bool) {
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

// GuardQuery is used for checking whethe a gin.Context contains valid token
func GuardQuery(c *gin.Context) (jwt.MapClaims, bool) {
	userToken := c.GetHeader("Authorization")
	trimmedTokenString := strings.TrimPrefix(userToken, `Bearer `)
	trimmedTokenString = strings.TrimSpace(trimmedTokenString)
	return isValidToken(trimmedTokenString)
}

// HandleVarifyToken is user for /api/veryfy_token
func HandleVarifyToken(c *gin.Context) {
	if tokenClaims, ok := GuardQuery(c); ok {
		fmt.Println("tokenClaims is ", tokenClaims)
		c.JSON(http.StatusOK, tokenClaims)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
