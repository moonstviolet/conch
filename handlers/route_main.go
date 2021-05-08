package handlers

import (
	"conch/data"
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
	session, err := data.CheckSession(r)
	t, _ := template.ParseFiles(
		"templates/index.html",
		"templates/lib/header.html",
		"templates/lib/question-flow.html",
	)
	if err != nil {
		qlist, _ := data.QuestionList()
		t.Execute(w, struct {
			User  data.User
			QList []data.Question
		}{
			QList: qlist,
		})
	} else {
		user := session.User()
		qlist, _ := data.QuestionList()
		t.Execute(w, struct {
			User  data.User
			QList []data.Question
		}{
			User:  user,
			QList: qlist,
		})
	}
}
