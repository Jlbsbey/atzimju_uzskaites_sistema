package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Data struct {
	IfMyself   bool      `json:"if_myself"`
	Username   string    `json:"username"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Role       string    `json:"role"`
	Email      string    `json:"email"`
	AvatarLink string    `json:"avatar_link"`
	Subjects   []Subject `json:"subjects"`
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	userID, err := strconv.Atoi(queryParams.Get("user"))
	originUserID := getUserID(session)
	if originUserID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if err != nil {
		userID = originUserID
	}
	sameUsers := userID == originUserID
	userData := getInfo(userID, sameUsers, originUserID)
	response := Response_Body{Status: "OK", Response: userData}
	json.NewEncoder(w).Encode(response)
}

func getInfo(userID int, sameUsers bool, adminID int) Data {
	var subjects []Subject
	var role string
	subjects, role = getSubjects(userID)
	var lg *sql.Rows
	var err error
	var name, surname, email, username, avatarLink string
	query := `SELECT username FROM login_details WHERE user_id = ?`
	lg, err = db.Query(query, userID)
	for lg.Next() {
		if err = lg.Scan(&username); err != nil {
			log.Println(err)
		}
	}
	switch role {
	case "student":
		query = `SELECT name, surname, email, avatar_link FROM students WHERE student_id = ?`
		lg, err = db.Query(query, userID)
	case "professor":
		query = `SELECT name, surname, email, avatar_link FROM professors WHERE professor_id = ?`
		lg, err = db.Query(query, userID)
	}
	for lg.Next() {
		if err = lg.Scan(&name, &surname, &email, &avatarLink); err != nil {
			log.Println(err)
		}
	}
	return Data{
		IfMyself:   sameUsers,
		Username:   username,
		Name:       name,
		Surname:    surname,
		Role:       role,
		Email:      email,
		Subjects:   subjects,
		AvatarLink: avatarLink,
	}

}
