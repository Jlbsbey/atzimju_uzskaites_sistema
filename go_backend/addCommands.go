package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

func AddSubject(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	name := queryParams.Get("subject")
	description := queryParams.Get("description")
	userID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if checkAdmin(userID) {
		addSubject(name, description)
		var response = Response_Body{Status: "OK"}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		var response = Response_Body{Status: "error", Error: "User does not have sufficient rights"}
		json.NewEncoder(w).Encode(response)
		return
	}
}

func checkAdmin(userID int) bool {
	var role string
	query := `SELECT role FROM login_details WHERE user_id = ?`
	lg, err := db.Query(query, userID)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&role); err != nil {
			log.Println(err)
		}
		if role == "admin" {
			return true
		}
	}
	return false
}

func addSubject(name string, description string) {
	exist := true
	var subjectID int
	for exist {
		subjectID = generateRandomInteger(1, math.MaxUint32)
		exist = checkExistance(subjectID)
	}
	query := `INSERT INTO subjects(subject_id, name, description) VALUES (?, ?, ?)`
	_, err := db.Query(query, subjectID, name, description)
	if err != nil {
		panic(err)
	}
}
