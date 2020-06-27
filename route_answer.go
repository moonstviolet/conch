package main

import (
	"conch/data"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func newAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, err := data.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			user := session.User()
			query := r.URL.Query()
			qid, _ := strconv.Atoi(query["qid"][0])
			question, _ := data.QuestionById(qid)
			t, _ := template.ParseFiles("templates/answer-new.html", "templates/lib/header.html")
			t.Execute(w, struct {
				U data.User
				Q data.Question
			}{
				U: user,
				Q: question,
			})
		}
	} else if r.Method == "POST" {
		session, err := data.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			user := session.User()
			query := r.URL.Query()
			qid, _ := strconv.Atoi(query["qid"][0])
			question, _ := data.QuestionById(qid)

			answer := data.Answer{
				Aid:      data.AutoIncrement("answers"),
				Qid:      question.Qid,
				Uid:      user.Uid,
				Username: user.Nickname,
				Detail:   r.PostFormValue("answerDetail"),
				Lastmod:  time.Now(),
			}
			answer.Create()
			question.ModUser = user.Nickname
			question.Modify = answer.Detail
			question.Lastmod = time.Now()
			err := question.Update()
			warning(err)
			http.Redirect(w, r, "/question/read?qid="+strconv.Itoa(question.Qid), http.StatusFound)
		}
	}
}

func readAnswer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	qid, _ := strconv.Atoi(query["qid"][0])
	question, _ := data.QuestionById(qid)
	t, _ := template.ParseFiles(
		"templates/question-read.html",
		"templates/lib/header.html",
		"templates/lib/question-profile.html",
	)
	session, _ := data.CheckSession(r)
	user := session.User()
	t.Execute(w, struct {
		U data.User
		Q data.Question
	}{
		U: user,
		Q: question,
	})
}
