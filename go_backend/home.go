package main

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")

	responseString := ""
	w.Write([]byte(responseString))
}
