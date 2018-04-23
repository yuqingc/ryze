package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
