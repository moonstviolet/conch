package handlers

import (
	"conch/models"
	"conch/proto"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func NewQuestion(c *gin.Context) {
	cookie, _ := c.Request.Cookie("session")
	session, _ := models.CheckSession(cookie.Value)
	user := session.User()

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "question-new.html", user)
		return
	}

	var req proto.NewQuestionReq
	if err := c.ShouldBind(&req); err != nil {
		c.Redirect(http.StatusFound, "/question/new")
		return
	}
	question := models.Question{
		Qid:      models.AutoIncrement("questions"),
		Uid:      user.Uid,
		Title:    req.QuestionTitle,
		Detail:   req.QuestionDetail,
		Follow:   1,
		Pageview: 1,
		Lastmod:  time.Now(),
	}
	if err := question.Create(); err != nil {
		log.Fatalf("Cannot create questionn, %v", err)
		c.Redirect(http.StatusFound, "/question/new")
		return
	}
	c.Redirect(http.StatusFound, "/question/read?qid="+strconv.Itoa(question.Qid))
}

func ReadQuestion(c *gin.Context) {
	var req proto.ReadQuestionReq
	if err := c.ShouldBind(&req); err != nil {
		//c.Redirect(http.StatusFound, "/error")
		return
	}

	question, _ := models.QuestionById(req.Qid)
	question.Pageview++
	question.Update()
	answers, _ := models.AnswersByQid(req.Qid)
	var user models.User
	cookie, err := c.Request.Cookie("session")
	if err == nil {
		if session, err := models.CheckSession(cookie.Value); err != nil {
			user = session.User()
		}
	}
	c.HTML(http.StatusOK, "question-read.html", gin.H{
		"LoginUser":  user,
		"Question":   question,
		"AnswerList": answers,
	})
}
