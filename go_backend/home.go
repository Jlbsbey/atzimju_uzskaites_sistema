package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Responce struct {
	LoginStatus string `json:"login_status"`
	Status      string `json:"status"`
	Error       string `json:"error"`
	Result      any    `json:"result"`
}

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
	SubjectID   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
}
type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type Response_Home struct {
	Role     string    `json:"role"`
	UserID   int       `json:"user_id"`
	Marks    []Mark    `json:"marks"`
	Subjects []Subject `json:"subjects"`
	Users    []User    `json:"users"` //IDs of professors if role = student and vice versa
}

type Response_Body struct {
	Status   string        `json:"status"`
	Error    string        `json:"error"`
	Response Response_Home `json:"content"`
}

var role string

func HomePage(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	userID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	subjects := getSubjects(userID)
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
	users := getUsers()
	if len(users) == 0 {
		var response = Response_Body{Status: "error", Error: "No students or professors"}
		json.NewEncoder(w).Encode(response)
		return
	}
	var response_home = Response_Home{Role: role, UserID: userID, Marks: Marks, Subjects: subjects, Users: users}
	var response = Response_Body{Status: "OK", Error: "", Response: response_home}
	json.NewEncoder(w).Encode(response)
}

func getUsers() []User {
	var users []User
	var lg *sql.Rows
	var ID int
	var err error
	var name, surname string
	if role == "student" {
		query := `SELECT professor_id, name, surname FROM professors `
		lg, err = db.Query(query)
		if err != nil {
			panic(err)
		}
	} else {
		query := `SELECT student_id, name, surname FROM students `
		lg, err = db.Query(query)
		if err != nil {
			panic(err)
		}
	}
	for lg.Next() {
		if err = lg.Scan(&ID, &name, &surname); err != nil {
			log.Println(err)
		}
		users = append(users, User{UserID: ID, Username: name + " " + surname})
	}
	return users
}

func getMarks(subject int, userID int) []Mark {
	var marks []Mark
	var number, markID, prof, studID int
	var create, edit time.Time
	query := `SELECT mark_id, student_id, professor_id, subject_id, value, create_date, edit_date FROM marks WHERE subject_id = ? AND student_id = ? OR professor_id = ?`
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
	}
	return marks

}

func getSubjects(userID int) []Subject {
	query := `SELECT subject_id FROM students_subjects WHERE student_id = ?`
	lg, err := db.Query(query, userID)
	var subjects []Subject
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		var ID int
		if err = lg.Scan(&ID); err != nil {
			log.Println(err)
		}
		query = `SELECT name FROM subjects WHERE subject_id = ?`
		lg, err = db.Query(query, ID)
		if err != nil {
			panic(err)
		}
		for lg.Next() {
			var name string
			if err = lg.Scan(&name); err != nil {
				log.Println(err)
			}
			subjects = append(subjects, Subject{SubjectID: ID, SubjectName: name})
		}
		role = "student"
	}
	if len(subjects) == 0 {
		query = `SELECT subject_id FROM professors_subjects WHERE professor_id = ?`
		lg, err = db.Query(query, userID)
		if err != nil {
			panic(err)
		}
		for lg.Next() {
			var ID int
			if err = lg.Scan(&ID); err != nil {
				log.Println(err)
			}
			query = `SELECT name FROM subjects WHERE subject_id = ?`
			lg, err = db.Query(query, userID)
			if err != nil {
				panic(err)
			}
			for lg.Next() {
				var name string
				if err = lg.Scan(&name); err != nil {
					log.Println(err)
				}
				subjects = append(subjects, Subject{SubjectID: ID, SubjectName: name})
			}
			role = "student"
		}
	}
	return subjects
}

func getUserID(session string) int {
	query := `SELECT user_id, expire_time FROM sessions WHERE session_key = ?`
	lg, err := db.Query(query, session)
	var DBexpire time.Time
	var login int
	if err != nil {
		print(1)
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&login, &DBexpire); err != nil {
			print(3)
			log.Println(err)
		}
		if time.Now().After(DBexpire) {
			return -1
		}
		return login
	}
	return -1
}
