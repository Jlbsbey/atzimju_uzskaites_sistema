package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func AddMark(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	studentID, _ := strconv.Atoi(queryParams.Get("student_id"))
	subjectID, _ := strconv.Atoi(queryParams.Get("subject"))
	mark, _ := strconv.Atoi(queryParams.Get("value"))
	markID, _ := strconv.Atoi(queryParams.Get("mark_id"))
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
	var formattedTime = time.Now().Format("2006-01-02 15:04:05")
	markID := generateRandomInteger(1000000000, 9999999999)
	query := `INSERT INTO marks(mark_id, student_id, professor_id, subject_id, value, create_date, edit_date) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.ExecContext(context.Background(), query, markID, studentID, profID, subjectID, mark, formattedTime, formattedTime)
	if err != nil {
		log.Println(err)
	}
}

func updateMark(markID int, mark int) {
	now := time.Now().Format("2006-01-02 15:04:05")
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
