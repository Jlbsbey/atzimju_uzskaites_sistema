package server

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
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

func ChangeUserSubjects(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	username := queryParams.Get("username")
	subjects := queryParams.Get("subjects")
	originUserID := getUserID(session)
	if originUserID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	userID := UserIDbyUsername(username)
	if checkAdmin(originUserID) {
		subjectList := strings.Split(subjects, ",")
		_, role := getSubjects(userID)
		err := updateSubjects(subjectList, userID, role)
		if err == 1 {
			var response = Response_Body{Status: "error", Error: "Incorrect role and/or username"}
			json.NewEncoder(w).Encode(response)
		}
		var response = Response_Body{Status: "OK"}
		json.NewEncoder(w).Encode(response)
	}
}

func updateSubjects(subjectList []string, userID int, role string) int {
	if role == "student" {
		query := `DELETE FROM students_subjects WHERE student_id = ?`
		_, err := db.Exec(query, userID)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(subjectList); i++ {
			intSubjectID, _ := strconv.Atoi(subjectList[i])
			query = `INSERT INTO students_subjects(pair_id, subject_id, student_id) VALUES (?, ?, ?)`
			pair_id := generateRandomInteger(1, math.MaxUint32)
			_, err = db.ExecContext(context.Background(), query, pair_id, intSubjectID, userID)
		}
		return 0
	} else if role == "professor" {
		query := `DELETE FROM professors_subjects WHERE professor_id = ?`
		_, err := db.Exec(query, userID)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(subjectList); i++ {
			intSubjectID, _ := strconv.Atoi(subjectList[i])
			query = `INSERT INTO professors_subjects(pair_id, subject_id, professor_id) VALUES (?, ?, ?)`
			pair_id := generateRandomInteger(1, math.MaxUint32)
			_, err = db.ExecContext(context.Background(), query, pair_id, intSubjectID, userID)
		}
		return 0
	}
	return 1
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
