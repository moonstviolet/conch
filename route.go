package main

import (
	"conch/data"
	"net/http"
	"text/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func index(w http.ResponseWriter, r *http.Request) {
	session, err := data.CheckSession(w, r)
	t, _ := template.ParseFiles("templates/index.html", "templates/lib/header.html")
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
			info(user.Username, "log in")
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
	t, _ := template.ParseFiles("templates/signup.html")
	t.Execute(w, nil)
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		danger(err)
	}
	user := data.User{
		Uuid:     primitive.NewObjectID(),
		Uid:      data.AutoIncrement("users"),
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
		Nickname: r.PostFormValue("nickname"),
	}
	if err := user.Create(); err != nil {
		danger(err)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
