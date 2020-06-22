package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/user/find", findUser)
	mux.HandleFunc("/user/profile", profile)

	mux.HandleFunc("/question/new", newQuestion)
	mux.HandleFunc("/question/read", readQuestion)

	mux.HandleFunc("/answer/new", newAnswer)

	server := http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
