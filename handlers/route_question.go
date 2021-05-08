package main

import (
	"conch/data"
	"net/http"
	"strconv"
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
			err = question.Create()
			if err != nil {
				danger(err)
			}
			http.Redirect(w, r, "/question/read?qid="+strconv.Itoa(question.Qid), http.StatusFound)
		}
	}
}

func readQuestion(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	qid, _ := strconv.Atoi(query["qid"][0])
	question, _ := data.QuestionById(qid)
	question.Pageview++
	question.Update()

	answers, _ := data.AnswersByQid(qid)

	session, _ := data.CheckSession(r)
	user := session.User()
	t, _ := template.ParseFiles(
		"templates/question-read.html",
		"templates/lib/header.html",
		"templates/lib/question-header.html",
		"templates/lib/answer-flow.html",
	)
	info(answers)
	t.Execute(w, struct {
		LoginUser  data.User
		Question   data.Question
		AnswerList []data.Answer
	}{
		LoginUser:  user,
		Question:   question,
		AnswerList: answers,
	})
}
