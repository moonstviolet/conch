package main

import (
	"conch/data"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signupAccount", signupAccount)
	server := http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}

func test() {
	user := data.User{
		Uuid:     primitive.NewObjectID(),
		Uid:      data.AutoIncrement("user"),
		Username: "moonstviolet",
		Password: "violet2943",
		Nickname: "Reverie",
	}
	user.Create()
}
