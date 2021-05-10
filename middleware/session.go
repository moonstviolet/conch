package middleware

import (
	"conch/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session")
		if err == nil {
			_, err = models.CheckSession(cookie.Value)
		}
		if err == nil {
			c.Next()
			return
		}
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
}
