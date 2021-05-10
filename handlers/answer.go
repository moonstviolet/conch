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

func NewAnswer(c *gin.Context) {
	cookie, _ := c.Cookie("session")
	session, _ := models.CheckSession(cookie)
	user := session.User()

	var req proto.NewAnswerReq
	if err := c.ShouldBind(&req); err != nil {
		//c.Redirect(http.StatusFound, "/error")
		return
	}
	question, _ := models.QuestionById(req.Qid)

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "answer-new.html", gin.H{
			"U": user,
			"Q": question,
		})
		return
	}

	answer := models.Answer{
		Aid:      models.AutoIncrement("answers"),
		Qid:      question.Qid,
		Uid:      user.Uid,
		Username: user.Nickname,
		Detail:   req.AnswerDetail,
		Lastmod:  time.Now(),
	}
	_ = answer.Create()
	question.ModUser = user.Nickname
	question.Modify = answer.Detail
	question.Lastmod = answer.Lastmod
	err := question.Update()
	if err != nil {
		log.Fatalf("Cannot create answer, %v", err)
	}
	c.Redirect(http.StatusFound, "/question/read?qid="+strconv.Itoa(question.Qid))
}

func ReadAnswer(c *gin.Context) {
	cookie, _ := c.Cookie("session")
	session, _ := models.CheckSession(cookie)
	user := session.User()

	var req proto.NewAnswerReq
	if err := c.ShouldBind(&req); err != nil {
		//c.Redirect(http.StatusFound, "/error")
		return
	}
	question, _ := models.QuestionById(req.Qid)
	c.HTML(http.StatusOK, "question-read.html", gin.H{
		"U": user,
		"Q": question,
	})
}
