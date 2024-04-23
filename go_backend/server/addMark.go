package server

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func AddMark(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	student_name := queryParams.Get("username")
	subjectID, _ := strconv.Atoi(queryParams.Get("subject_id"))
	mark, _ := strconv.Atoi(queryParams.Get("value"))
	markID, _ := strconv.Atoi(queryParams.Get("mark_id"))
	if mark > 10 || mark <= 0 {
		var response = Response_Body{Status: "error", Error: "Wrong mark. Marks are only from 1 to 10"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	studentID := UserIDbyUsername(student_name)
	userID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if markID == 0 {
		insertMark(userID, studentID, subjectID, mark)
		response := Response_Body{Status: "OK"}
		json.NewEncoder(w).Encode(response)
		return
	}
	updateMark(markID, mark)
	response := Response_Body{Status: "OK"}
	json.NewEncoder(w).Encode(response)
}

func insertMark(profID int, studentID int, subjectID int, mark int) {
	formattedTime := time.Now().In(location).Format("2006-01-02 15:04:05")
	exist := true
	var markID int
	for exist {
		markID = generateRandomInteger(1, math.MaxUint32)
		exist = checkExistance(markID)
	}
	query := `INSERT INTO marks(mark_id, student_id, professor_id, subject_id, value, create_date, edit_date) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.ExecContext(context.Background(), query, markID, studentID, profID, subjectID, mark, formattedTime, formattedTime)
	if err != nil {
		log.Println(err)
	}
}

func checkExistance(markID int) bool {
	query := `SELECT mark_id FROM marks WHERE mark_id = ?`
	lg, err := db.Query(query, markID)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		return true
	}
	return false
}

func updateMark(markID int, mark int) {
	now := time.Now().In(location).Format("2006-01-02 15:04:05")
	query := `UPDATE marks SET value = ?, edit_date = ? WHERE mark_id = ?`
	_, err := db.Query(query, mark, now, markID)
	if err != nil {
		panic(err)
	}
}

func generateRandomInteger(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
