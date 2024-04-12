package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Student struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

func StudentList(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	name := queryParams.Get("name")
	userID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	students := getStudents(name)
	response := Response_Body{Status: "OK", Response: students}
	json.NewEncoder(w).Encode(response)
}

func getStudents(part string) []Student {
	var name, surname, email string
	var students []Student
	query := `SELECT name, surname, email FROM students WHERE name LIKE ? OR surname LIKE ?`
	lg, err := db.Query(query, "%"+part+"%", "%"+part+"%")
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&name, &surname, &email); err != nil {
			log.Println(err)
		}
		students = append(students, Student{Name: name, Surname: surname, Email: email})
	}
	return students
}
