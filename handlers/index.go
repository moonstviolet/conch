package handlers

import (
	"conch/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	qlist, _ := models.QuestionList()
	cookie, err := c.Request.Cookie("session")
	if err == nil {
		if session, err := models.CheckSession(cookie.Value); err == nil {
			user := session.User()
			c.HTML(http.StatusOK, "index.html", gin.H{
				"QList": qlist,
				"User":  user,
			})
			return
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"QList": qlist,
	})
}
