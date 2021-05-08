package main

import (
	"conch/data"
	"encoding/json"
	"net/http"
	"text/template"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		user, err := data.UserByUsername(r.PostFormValue("username"))
		if err != nil || user.Password != r.PostFormValue("password") {
			// TODO: 最好能提示用户名或密码错误
			http.Redirect(w, r, "login", http.StatusFound)
		} else {
			session, err := user.CreateSession()
			if err != nil {
				danger(err, "Cannot create session")
			}
			cookie := http.Cookie{
				Name:     "session",
				Value:    session.Sid,
				HttpOnly: true,
				MaxAge:   3600,
			}
			http.SetCookie(w, &cookie)
			info("user", user.Username, "login")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else {
		if _, err := data.CheckSession(r); err != nil {
			t, _ := template.ParseFiles("templates/login.html", "templates/lib/header.html")
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != http.ErrNoCookie {
		session := data.Session{
			Sid: cookie.Value,
		}
		session.DeleteBySid()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			danger(err)
		}
		user := data.User{
			Uid:      data.AutoIncrement("users"),
			Username: r.PostFormValue("username"),
			Password: r.PostFormValue("password"),
			Email:    r.PostFormValue("email"),
			Nickname: r.PostFormValue("nickname"),
			Motto:    r.PostFormValue("motto"),
		}
		user.EmailHash()
		if err := user.Create(); err != nil {
			danger(err)
		} else {
			// TODO: 最好能提示注册成功
			info("user", user.Username, "signup")
			http.Redirect(w, r, "login", http.StatusFound)
		}
	} else {
		if _, err := data.CheckSession(r); err != nil {
			t, _ := template.ParseFiles("templates/signup.html", "templates/lib/header.html")
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func findUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if _, err := data.UserByUsername(query["username"][0]); err != nil {
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

func profile(w http.ResponseWriter, r *http.Request) {
	session, err := data.CheckSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		user := session.User()
		answers, _ := data.AnswersByUid(user.Uid)
		t, _ := template.ParseFiles(
			"templates/user-profile.html", 
			"templates/lib/header.html",
			"templates/lib/answer-flow.html",
		)
		t.Execute(w, struct {
			LoginUser  data.User
			AnswerList []data.Answer
		}{
			LoginUser:  user,
			AnswerList: answers,
		})
	}
}
