package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Mark struct {
	number     byte
	studentID  string
	profID     string
	createDate string
	editDate   string
	markID     string
	subject    string
}

var role string

func HomePage(w http.ResponseWriter, r *http.Request /*session string*/) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	fmt.Println(session)
	userID := getUserID(session)
	if userID == "" {
		w.Write([]byte("auth:404")) //истекло время сессии или пользователь не был найден по сессии
		//print("wrong userid")
		//return
	}
	subjects := getSubjects(userID)
	if len(subjects) == 0 {
		w.Write([]byte("auth:405")) //предметов нет у этого человека, выводи сообщение что бы админ их добавил
		//print("no subjects")
		//return
	}
	var response []Mark
	for i := 0; i < len(subjects); i++ {
		markarr := getMarks(subjects[i], role, userID)
		for j := 0; j < len(markarr); j++ {
			response = append(response, markarr[j])
		}
	}
	/*for i := 0; i < len(response); i++ {
		println("Mark ", i)
		println("MarkID", response[i].markID)
		println("StudID", response[i].studentID)
		println("ProfID", response[i].profID)
		println("CreateDate", response[i].createDate)
		println("EditDate", response[i].editDate)
		println("subject", response[i].subject)
		println("number", response[i].number)
		println("___________________________")
	}*/
	responseString, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Write(responseString)
}

func getMarks(subject string, role string, userID string) []Mark {
	var marks []Mark
	var number byte
	var markID string
	var create, edit time.Time
	if role == "stud" {
		query := `SELECT markID, profID, mark, createDate, editDate FROM marks WHERE subjID = ? AND studID = ?`
		lg, err := db.Query(query, subject, userID)
		if err != nil {
			panic(err)
		}
		var prof string
		for lg.Next() {
			if err = lg.Scan(&markID, &prof, &number, &create, &edit); err != nil {
				log.Println(err)
			}
			createTimeformat := create.Format("2006-01-02 15:04:05")
			editTimeFormat := create.Format("2006-01-02 15:04:05")
			marks = append(marks, Mark{markID: markID, studentID: "", profID: prof, number: number, createDate: createTimeformat, editDate: editTimeFormat, subject: subject})
		}
		return marks
	}
	query := `SELECT markID, studID, mark, createDate, editDate FROM marks WHERE subjID = ? AND profID = ?`
	lg, err := db.Query(query, subject, userID)
	if err != nil {
		panic(err)
	}
	var stud string
	for lg.Next() {
		if err = lg.Scan(&markID, &stud, &number, &create, &edit); err != nil {
			log.Println(err)
		}

		createTimeformat := create.Format("2006-01-02 15:04:05")
		editTimeFormat := create.Format("2006-01-02 15:04:05")
		marks = append(marks, Mark{markID: markID, studentID: stud, profID: "", number: number, createDate: createTimeformat, editDate: editTimeFormat, subject: subject})
	}
	return marks

}

func getSubjects(userID string) []string {
	query := `SELECT subjID FROM studsubj WHERE studID = ?`
	lg, err := db.Query(query, userID)
	var subjects []string
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		var temp string
		if err = lg.Scan(&temp); err != nil {
			log.Println(err)
		}
		subjects = append(subjects, temp)
		role = "stud"
	}
	if len(subjects) == 0 {
		query = `SELECT subjID FROM profsubj WHERE profID = ?`
		lg, err = db.Query(query, userID)
		if err != nil {
			panic(err)
		}
		for lg.Next() {
			var temp string
			if err = lg.Scan(&temp); err != nil {
				log.Println(err)
			}
			subjects = append(subjects, temp)
			role = "prof"
		}
	}
	return subjects
}

func getUserID(session string) string {
	query := `SELECT login, expire FROM sessions WHERE sessionID = ?`
	lg, err := db.Query(query, session)
	var DBexpire time.Time
	var login string
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&login, &DBexpire); err != nil {
			log.Println(err)
		}
		deltaTime := DBexpire.Sub(time.Now())
		if deltaTime <= 0 {
			return ""
		}
		return login
	}
	return ""
}
