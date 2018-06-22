package blogs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuqingc/ryze/pkg/auth"
)

// HandlePostBlog : create a blog
func HandlePostBlog(c *gin.Context) {
	if _, ok := auth.GuardQuery(c); ok {
		c.String(http.StatusOK, "ok")
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
