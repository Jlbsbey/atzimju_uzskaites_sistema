package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
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

func ChangeUserData(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	userID, _ := strconv.Atoi(queryParams.Get("user"))
	email := queryParams.Get("email")
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	originUserID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if checkAdmin(originUserID) {
		if email != "" {
			updateEmail(userID, email)
		}
		if name != "" {
			updateName(userID, name)
		}
		if surname != "" {
			updateSurname(userID, surname)
		}
		var response = Response_Body{Status: "OK"}
		json.NewEncoder(w).Encode(response)
		return
	}
}

func ChangeUserSubjects(w http.ResponseWriter, r *http.Request) {
}

func getSubjectsName(part string) []Subject {
	var name, description string
	var ID int
	var students []Subject
	query := `SELECT * FROM subjects WHERE name LIKE ? `
	lg, err := db.Query(query, "%"+part+"%")
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&ID, &name, &description); err != nil {
			log.Println(err)
		}
		students = append(students, Subject{SubjectID: ID, SubjectName: name, SubjectDescription: description})
	}
	return students
}

func SearchSubjects(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	name := queryParams.Get("subject")
	userID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	subjects := getSubjectsName(name)
	response := Response_Body{Status: "OK", Response: subjects}
	json.NewEncoder(w).Encode(response)
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
