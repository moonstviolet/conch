package handlers

import (
	"conch/error_code"
	"conch/models"
	"conch/proto"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func NewQuestion(rep *proto.NewQuestionReq, resp *proto.NewQuestionResp) *error_code.RespError {
	if r.Method == "GET" {
		session, err := models.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			user := session.User()
			t, _ := template.ParseFiles("templates/question-new.html", "templates/lib/header.html")
			t.Execute(w, user)
		}
	} else if r.Method == "POST" {
		session, err := models.CheckSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			err = r.ParseForm()
			if err != nil {
				//danger(err)
			}
			user := session.User()
			question := models.Question{
				Qid:      models.AutoIncrement("questions"),
				Uid:      user.Uid,
				Title:    r.PostFormValue("questionTitle"),
				Detail:   r.PostFormValue("questionDetail"),
				Follow:   1,
				Pageview: 1,
				Lastmod:  time.Now(),
			}
			err = question.Create()
			if err != nil {
				//danger(err)
			}
			http.Redirect(w, r, "/question/read?qid="+strconv.Itoa(question.Qid), http.StatusFound)
		}
	}
}

func ReadQuestion(rep *proto.ReadQuestionReq, resp *proto.ReadQuestionResp) *error_code.RespError {
	query := r.URL.Query()
	qid, _ := strconv.Atoi(query["qid"][0])
	question, _ := models.QuestionById(qid)
	question.Pageview++
	question.Update()

	answers, _ := models.AnswersByQid(qid)

	session, _ := models.CheckSession(r)
	user := session.User()
	t, _ := template.ParseFiles(
		"templates/question-read.html",
		"templates/lib/header.html",
		"templates/lib/question-header.html",
		"templates/lib/answer-flow.html",
	)
	t.Execute(w, struct {
		LoginUser  models.User
		Question   models.Question
		AnswerList []models.Answer
	}{
		LoginUser:  user,
		Question:   question,
		AnswerList: answers,
	})
}
