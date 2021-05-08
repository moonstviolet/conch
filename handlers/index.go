package handlers

import (
	"conch/error_code"
	"conch/models"
	"conch/proto"
	"text/template"
)

func Index(rep *proto.IndexReq, resp *proto.IndexResp) *error_code.RespError {
	session, err := models.CheckSession(rep)
	t, _ := template.ParseFiles(
		"templates/index.html",
		"templates/lib/header.html",
		"templates/lib/question-flow.html",
	)
	if err != nil {
		qlist, _ := models.QuestionList()
		t.Execute(w, struct {
			User  models.User
			QList []models.Question
		}{
			QList: qlist,
		})
	} else {
		user := session.User()
		qlist, _ := models.QuestionList()
		t.Execute(w, struct {
			User  models.User
			QList []models.Question
		}{
			User:  user,
			QList: qlist,
		})
	}
	return nil
}
