package handlers

import (
	"conch/models"
	"net/http"
	"text/template"
)

func cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h(w, r)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	session, err := models.CheckSession(r)
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
}
