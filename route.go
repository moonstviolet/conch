package main

import (
	"conch/data"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	session, err := data.CheckSession(w, r)
	t, _ := template.ParseFiles(
		"templates/index.html",
		"templates/lib/header.html",
		"templates/lib/question.html",
		"templates/lib/question-flow.html",
	)
	if err != nil {
		t.Execute(w, nil)
	} else {
		user := session.User()
		t.Execute(w, user)
	}
}

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
			}
			http.SetCookie(w, &cookie)
			info("user", user.Username, "login")
			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else {
		if _, err := data.CheckSession(w, r); err != nil {
			t, _ := template.ParseFiles("templates/login.html", "templates/lib/header.html")
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
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
			Nickname: r.PostFormValue("nickname"),
			Motto:    r.PostFormValue("motto"),
		}
		if err := user.Create(); err != nil {
			danger(err)
		} else {
			// TODO: 最好能提示注册成功
			info("user", user.Username, "signup")
			http.Redirect(w, r, "login", http.StatusFound)
		}
	} else {
		if _, err := data.CheckSession(w, r); err != nil {
			t, _ := template.ParseFiles("templates/signup.html", "templates/lib/header.html")
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func finduser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if _, err := data.UserByUsername(query["username"][0]); err != nil {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func ask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	} else {
		t, _ := template.ParseFiles("templates/ask.html", "templates/lib/header.html")
		t.Execute(w, nil)
	}
}
