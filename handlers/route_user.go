package handlers

import (
	"conch/models"
	"encoding/json"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		user, err := models.UserByUsername(r.PostFormValue("username"))
		if err != nil || user.Password != r.PostFormValue("password") {
			// TODO: 最好能提示用户名或密码错误
			http.Redirect(w, r, "login", http.StatusFound)
		} else {
			session, err := user.CreateSession()
			if err != nil {
				//danger(err, "Cannot create session")
			}
			cookie := http.Cookie{
				Name:     "session",
				Value:    session.Sid,
				HttpOnly: true,
				MaxAge:   3600,
			}
			http.SetCookie(w, &cookie)
			//info("user", user.Username, "login")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else {
		if _, err := models.CheckSession(r); err != nil {
			t, _ := template.ParseFiles("templates/login.html", "templates/lib/header.html")
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != http.ErrNoCookie {
		session := models.Session{
			Sid: cookie.Value,
		}
		session.DeleteBySid()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			//danger(err)
		}
		user := models.User{
			Uid:      models.AutoIncrement("users"),
			Username: r.PostFormValue("username"),
			Password: r.PostFormValue("password"),
			Email:    r.PostFormValue("email"),
			Nickname: r.PostFormValue("nickname"),
			Motto:    r.PostFormValue("motto"),
		}
		user.EmailHash()
		if err := user.Create(); err != nil {
			//danger(err)
		} else {
			// TODO: 最好能提示注册成功
			//info("user", user.Username, "signup")
			http.Redirect(w, r, "login", http.StatusFound)
		}
	} else {
		if _, err := models.CheckSession(r); err != nil {
			t, _ := template.ParseFiles("templates/signup.html", "templates/lib/header.html")
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if _, err := models.UserByUsername(query["username"][0]); err != nil {
		b, _ := json.Marshal(struct {
			IsValid bool `json:"isValid"`
		}{
			IsValid: true,
		})
		w.Write(b)
	} else {
		b, _ := json.Marshal(struct {
			IsValid bool `json:"isValid"`
		}{
			IsValid: false,
		})
		w.Write(b)
	}
}

func Profile(w http.ResponseWriter, r *http.Request) {
	session, err := models.CheckSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		user := session.User()
		answers, _ := models.AnswersByUid(user.Uid)
		t, _ := template.ParseFiles(
			"templates/user-profile.html",
			"templates/lib/header.html",
			"templates/lib/answer-flow.html",
		)
		t.Execute(w, struct {
			LoginUser  models.User
			AnswerList []models.Answer
		}{
			LoginUser:  user,
			AnswerList: answers,
		})
	}
}
