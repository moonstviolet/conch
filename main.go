package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signupAccount", signupAccount)
	server := http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	info("server start")
	server.ListenAndServe()
}
