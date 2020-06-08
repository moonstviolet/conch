package main

import (
	"conch/data"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Address string `json:"Address"`
	Static  string `json:"Static"`
}

var config Configuration
var logger *log.Logger

func init() {
	loadConfig()
	file, err := os.OpenFile("conch.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

func checkSession(w http.ResponseWriter, r *http.Request) (session data.Session, err error) {
	cookie, err := r.Cookie("session")
	if err == nil {
		session = data.Session{Sid: cookie.Value}
		if !session.Check() {
			err = errors.New("Invalid session")
		}
	}
	return
}
