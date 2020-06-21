package main

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

func index(w http.ResponseWriter, r *http.Request) {
	session, err := data.CheckSession(r)
	t, _ := template.ParseFiles(
		"templates/index.html",
		"templates/lib/header.html",
		"templates/lib/question-flow.html",
	)
	if err != nil {
		t.Execute(w, nil)
	} else {
		user := session.User()
		question, _ := data.QuestionById(18)
		t.Execute(w, struct {
			U data.User
			Q data.Question
		}{
			U: user,
			Q: question,
		})
	}
}
