package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("Something went wrong. Please try again"))
		return false
	}

	isvalid := (username == USERNAME) && (password == PASSWORD)
	if !isvalid {
		w.Write([]byte("Invalid username or password"))
		return false
	}
	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Invalid method, only GET requests are allowed"))
		return false
	}

	return true
}
