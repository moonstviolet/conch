package handlers

import (
	"conch/models"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func NewAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, err := models.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			user := session.User()
			query := r.URL.Query()
			qid, _ := strconv.Atoi(query["qid"][0])
			question, _ := models.QuestionById(qid)
			t, _ := template.ParseFiles("templates/answer-new.html", "templates/lib/header.html")
			t.Execute(w, struct {
				U models.User
				Q models.Question
			}{
				U: user,
				Q: question,
			})
		}
	} else if r.Method == "POST" {
		session, err := models.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			user := session.User()
			query := r.URL.Query()
			qid, _ := strconv.Atoi(query["qid"][0])
			question, _ := models.QuestionById(qid)

			answer := models.Answer{
				Aid:      models.AutoIncrement("answers"),
				Qid:      question.Qid,
				Uid:      user.Uid,
				Username: user.Nickname,
				Detail:   r.PostFormValue("answerDetail"),
				Lastmod:  time.Now(),
			}
			answer.Create()
			question.ModUser = user.Nickname
			question.Modify = answer.Detail
			question.Lastmod = answer.Lastmod
			err := question.Update()
			if err != nil {
				//
			}
			http.Redirect(w, r, "/question/read?qid="+strconv.Itoa(question.Qid), http.StatusFound)
		}
	}
}

func ReadAnswer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	qid, _ := strconv.Atoi(query["qid"][0])
	question, _ := models.QuestionById(qid)
	t, _ := template.ParseFiles(
		"templates/question-read.html",
		"templates/lib/header.html",
		"templates/lib/question-profile.html",
	)
	session, _ := models.CheckSession(r)
	user := session.User()
	t.Execute(w, struct {
		U models.User
		Q models.Question
	}{
		U: user,
		Q: question,
	})
}
