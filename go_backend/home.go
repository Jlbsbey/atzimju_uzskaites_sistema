package main

import (
	"encoding/json"
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

type Response_Home struct {
	Role   string `json:"role"`
	UserID int    `json:"user_id"`
	Marks  []Mark `json:"marks"`
}

/*export interface Grade {
	id: number;
	subject_id: number;
	value: number;
	professor: string;
	created_unix: number;
	last_updated_unix: number;
}*/

var role string

func HomePage(w http.ResponseWriter, r *http.Request /*session string*/) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	userID := getUserID(session)
	if userID == -1 {
		w.Write([]byte("auth:404")) //истекло время сессии или пользователь не был найден по сессии
		return
	}
	subjects := getSubjects(userID)
	if len(subjects) == 0 {
		w.Write([]byte("auth:405")) //предметов нет у этого человека, выводи сообщение что бы админ их добавил
		return
	}
	var Marks []Mark
	for i := 0; i < len(subjects); i++ {
		markarr := getMarks(subjects[i], userID)
		for j := 0; j < len(markarr); j++ {
			Marks = append(Marks, markarr[j])
		}
	}
	var response = Response_Home{Role: role, UserID: userID, Marks: Marks}
	json.NewEncoder(w).Encode(response)
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

func getSubjects(userID int) []int {
	query := `SELECT subject_id FROM students_subjects WHERE student_id = ?`
	lg, err := db.Query(query, userID)
	var subjects []int
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		var temp int
		if err = lg.Scan(&temp); err != nil {
			log.Println(err)
		}
		subjects = append(subjects, temp)
		role = "student"
	}
	if len(subjects) == 0 {
		query = `SELECT subject_id FROM professors_subjects WHERE professor_id = ?`
		lg, err = db.Query(query, userID)
		if err != nil {
			panic(err)
		}
		for lg.Next() {
			var temp int
			if err = lg.Scan(&temp); err != nil {
				log.Println(err)
			}
			subjects = append(subjects, temp)
			role = "professor"
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
		print(2)
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
