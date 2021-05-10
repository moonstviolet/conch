package handlers

import (
	"conch/models"
	"conch/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		cookie, err := c.Request.Cookie("session")
		if err == nil {
			_, err = models.CheckSession(cookie.Value)
		}
		if err != nil {
			c.HTML(http.StatusOK, "login.html", "")
		} else {
			c.Redirect(http.StatusFound, "/")
		}
		return
	}

	var req proto.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	user, err := models.UserByUsername(req.Username)
	if err != nil || user.Password != req.Password {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	session, err := user.CreateSession()
	if err != nil {
		log.Fatalf("Cannot create session, %v", err)
	}
	c.SetCookie("session", session.Sid, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
}

func Logout(c *gin.Context) {
	cookie, _ := c.Request.Cookie("session")
	session, _ := models.CheckSession(cookie.Value)
	_ = session.DeleteBySid()
	c.Redirect(http.StatusFound, "/")
}

func Signup(c *gin.Context) {
	if c.Request.Method == "GET" {
		cookie, err := c.Request.Cookie("session")
		if err == nil {
			_, err = models.CheckSession(cookie.Value)
		}
		if err != nil {
			c.HTML(http.StatusOK, "signup.html", "")
		} else {
			c.Redirect(http.StatusFound, "/")
		}
		return
	}

	var req proto.SignupReq
	if err := c.ShouldBind(&req.User); err != nil {
		c.Redirect(http.StatusFound, "/signup")
		return
	}
	req.User.Uid = models.AutoIncrement("users")
	req.User.EmailHash()
	if err := req.User.Create(); err != nil {
		// todo
		log.Fatalf("Cannot create user, %v", err)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

// func FindUser(rep *proto.FindUserReq, resp *proto.FindUserResp) *error_code.RespError {
// 	query := r.URL.Query()
// 	if _, err := models.UserByUsername(query["username"][0]); err != nil {
// 		b, _ := json.Marshal(struct {
// 			IsValid bool `json:"isValid"`
// 		}{
// 			IsValid: true,
// 		})
// 		w.Write(b)
// 	} else {
// 		b, _ := json.Marshal(struct {
// 			IsValid bool `json:"isValid"`
// 		}{
// 			IsValid: false,
// 		})
// 		w.Write(b)
// 	}
// }

func Profile(c *gin.Context) {
	cookie, _ := c.Request.Cookie("session")
	session, _ := models.CheckSession(cookie.Value)
	user := session.User()
	answers, _ := models.AnswersByUid(user.Uid)
	c.HTML(http.StatusOK, "user-profile.html", gin.H{
		"LoginUser":  user,
		"AnswerList": answers,
	})
}
