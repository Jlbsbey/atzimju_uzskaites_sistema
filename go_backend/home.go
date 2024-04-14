package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Mark struct {
	Number     int    `json:"value"`
	StudentID  int    `json:"student_id"`
	ProfID     int    `json:"professor_id"`
	CreateDate string `json:"create_date"`
	EditDate   string `json:"edit_date"`
	MarkID     int    `json:"mark_id"`
	SubjectID  int    `json:"subject_id"`
}
type Subject struct {
	SubjectID          int    `json:"subject_id"`
	SubjectName        string `json:"subject_name"`
	SubjectDescription string `json:"subject_description"`
}
type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"name"`
}

type Response_Home struct {
	Role       string    `json:"role"`
	UserID     int       `json:"user_id"`
	Marks      []Mark    `json:"marks"`
	Subjects   []Subject `json:"subjects"`
	Students   []User    `json:"students"`
	Professors []User    `json:"professors"`
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	userID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	subjects, role := getSubjects(userID)
	if len(subjects) == 0 {
		var response = Response_Body{Status: "error", Error: "No subjects"} //предметов нет у этого человека, выводи сообщение что бы админ их добавил
		json.NewEncoder(w).Encode(response)
		return
	}
	var Marks []Mark
	for i := 0; i < len(subjects); i++ {
		markarr := getMarks(subjects[i].SubjectID, userID)
		for j := 0; j < len(markarr); j++ {
			Marks = append(Marks, markarr[j])
		}
	}
	students, professors := getUsers()
	if len(students) == 0 {
		var response = Response_Body{Status: "error", Error: "No students"}
		json.NewEncoder(w).Encode(response)
		return
	}
	if len(professors) == 0 {
		var response = Response_Body{Status: "error", Error: "No professors"}
		json.NewEncoder(w).Encode(response)
		return
	}
	var response_home = Response_Home{Role: role, UserID: userID, Marks: Marks, Subjects: subjects, Students: students, Professors: professors}
	var response = Response_Body{Status: "OK", Error: "", Response: response_home}
	json.NewEncoder(w).Encode(response)
}

func getUsers() ([]User, []User) {
	var students, professors []User
	var lg *sql.Rows
	var ID int
	var err error
	var name, surname string
	query := `SELECT professor_id, name, surname FROM professors `
	lg, err = db.Query(query)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&ID, &name, &surname); err != nil {
			log.Println(err)
		}
		professors = append(professors, User{UserID: ID, Username: name + " " + surname})
	}
	query = `SELECT student_id, name, surname FROM students `
	lg, err = db.Query(query)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&ID, &name, &surname); err != nil {
			log.Println(err)
		}
		students = append(students, User{UserID: ID, Username: name + " " + surname})
	}
	return students, professors
}

func getMarks(subject int, userID int) []Mark {
	var marks []Mark
	var number, markID, prof, studID int
	var create, edit time.Time
	query := `SELECT mark_id, student_id, professor_id, subject_id, value, create_date, edit_date FROM marks WHERE subject_id = ? AND (student_id = ? OR professor_id = ?)`

	lg, err := db.Query(query, subject, userID, userID)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&markID, &studID, &prof, &subject, &number, &create, &edit); err != nil {
			log.Println(err)
		}
		createTimeformat := create.Format("2006-01-02 15:04:05")
		editTimeFormat := edit.Format("2006-01-02 15:04:05")
		marks = append(marks, Mark{MarkID: markID, StudentID: studID, ProfID: prof, Number: number, CreateDate: createTimeformat, EditDate: editTimeFormat, SubjectID: subject})
		fmt.Println(len(marks))
	}
	return marks

}

func getSubjects(userID int) ([]Subject, string) {
	var role string
	query := `SELECT role FROM login_details WHERE user_id = ?`
	lg, err := db.Query(query, userID)
	for lg.Next() {
		if err = lg.Scan(&role); err != nil {
			log.Println(err)
		}
	}
	var subjects []Subject
	if role == "student" {
		query = `SELECT subject_id FROM students_subjects WHERE student_id = ?`
		lg, err = db.Query(query, userID)
		if err != nil {
			panic(err)
		}
	} else if role == "professor" {
		query = `SELECT subject_id FROM professors_subjects WHERE professor_id = ?`
		lg, err = db.Query(query, userID)
		if err != nil {
			panic(err)
		}
	} else if role == "admin" {
		var subjID int
		var name, description string
		query = `SELECT * FROM subjects`
		lg, err = db.Query(query)
		if err != nil {
			panic(err)
		}
		for lg.Next() {
			if err = lg.Scan(&subjID, &name, &description); err != nil {
				log.Println(err)
			}
			subjects = append(subjects,
				Subject{
					SubjectID:          subjID,
					SubjectName:        name,
					SubjectDescription: description,
				},
			)
		}
	}
	for lg.Next() {
		var ID int
		var innerLg *sql.Rows
		if err = lg.Scan(&ID); err != nil {
			log.Println(err)
		}
		query = `SELECT name, description FROM subjects WHERE subject_id = ?`
		innerLg, err = db.Query(query, ID)
		if err != nil {
			panic(err)
		}
		for innerLg.Next() {
			var name, description string
			if err = innerLg.Scan(&name, &description); err != nil {
				log.Println(err)
			}
			subjects = append(subjects,
				Subject{
					SubjectID:          ID,
					SubjectName:        name,
					SubjectDescription: description,
				},
			)
		}
	}
	return subjects, role
}

func getUserID(session string) int {
	ClearSessionsOnce()
	query := `SELECT user_id, expire_time FROM sessions WHERE session_key = ?`
	lg, err := db.Query(query, session)
	var DBexpire time.Time
	var login int
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&login, &DBexpire); err != nil {
			log.Println(err)
		}
		if time.Now().After(DBexpire) {
			return -1
		}
		return login
	}
	return -1
}
