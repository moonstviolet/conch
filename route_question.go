package main

import (
	"conch/data"
	"net/http"
	"text/template"
	"time"
)

func newQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, err := data.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			user := session.User()
			t, _ := template.ParseFiles("templates/question-new.html", "templates/lib/header.html")
			t.Execute(w, user)
		}
	} else if r.Method == "POST" {
		session, err := data.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			err = r.ParseForm()
			if err != nil {
				danger(err)
			}
			user := session.User()
			question := data.Question{
				Qid:      data.AutoIncrement("questions"),
				Uid:      user.Uid,
				Title:    r.PostFormValue("questionTitle"),
				Detail:   r.PostFormValue("questionDetail"),
				Follow:   1,
				Pageview: 1,
				Lastmod:  time.Now(),
			}
			question.Create()
			http.Redirect(w, r, "/question/read", http.StatusFound)
		}
	}
}

func readQuestion(w http.ResponseWriter, r *http.Request) {
	session, err := data.CheckSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		user := session.User()
		question, _ := data.QuestionById(18)
		t, _ := template.ParseFiles(
			"templates/question-read.html",
			"templates/lib/header.html",
			"templates/lib/question-profile.html",
		)
		t.Execute(w, struct {
			U data.User
			Q data.Question
		}{
			U: user,
			Q: question,
		})
	}
}
