package handlers

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html,header.html", "")

	// t, _ := template.ParseFiles(
	// 	"templates/index.html",
	// 	"templates/lib/header.html",
	// 	"templates/lib/question-flow.html",
	// )
	// if err != nil {
	// 	qlist, _ := models.QuestionList()
	// 	t.Execute(w, struct {
	// 		User  models.User
	// 		QList []models.Question
	// 	}{
	// 		QList: qlist,
	// 	})
	// } else {
	// 	user := session.User()
	// 	qlist, _ := models.QuestionList()
	// 	t.Execute(w, struct {
	// 		User  models.User
	// 		QList []models.Question
	// 	}{
	// 		User:  user,
	// 		QList: qlist,
	// 	})
	// }
	// return nil
}
