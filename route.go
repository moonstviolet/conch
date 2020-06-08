package main

import (
	"conch/data"
	"net/http"
	"text/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func index(w http.ResponseWriter, r *http.Request) {
	session, err := checkSession(w, r)
	t, _ := template.ParseFiles("templates/index.html", "templates/packageHeader.html")
	if err != nil {
		user := data.User{
			Nickname: "Reverie",
			Uid:      1,
		}
		t.Execute(w, user)
	} else {
		user := session.User()
		t.Execute(w, user)
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
		info("signup succeed")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
