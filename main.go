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
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/finduser", finduser)
	mux.HandleFunc("/ask", ask)
	server := http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
