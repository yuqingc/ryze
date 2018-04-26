package auth

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

const testToken string = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjQ4MzAzMzksImlhdCI6MTUyNDc0MzkzOSwiaXNzIjoiUnl6ZS9nZW5lcmFsIiwidXNyIjoiYWRtaW4ifQ.BqF34UJOm37DsXitIDJ5MSs_VXsZpu8D0gS9pSL9dcc`

// IsValidToken checks if a token is valid
func IsValidToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my-secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return true
	}
	fmt.Println(err)
	return false
}

func init() {
	fmt.Println("init, ", IsValidToken(testToken))
}
