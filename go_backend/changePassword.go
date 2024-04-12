package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	email := queryParams.Get("email")
	oldPassword := queryParams.Get("oldPassword")
	newPassword := queryParams.Get("newPassword")
	userID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if email != "" && oldPassword == "" {
		updateEmail(userID, email)
		var response = Response_Body{Status: "OK", Error: ""} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if checkValidity(userID, oldPassword) {
		if email != "" {
			updateEmail(userID, email)
		}
		updatePassword(userID, newPassword)
		//clearUserSessions(userID)
		var response = Response_Body{Status: "OK", Error: ""} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	//очистка всех сессий связанных с юзерид
	var response = Response_Body{Status: "error", Error: "Password is incorrect"} //истекло время сессии или пользователь не был найден по сессии
	json.NewEncoder(w).Encode(response)
}

func checkValidity(userID int, oldPassword string) bool {
	var password, salt string
	query := `SELECT password, password_salt FROM login_details WHERE user_id = ?`
	lg, err := db.Query(query, userID)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&password, &salt); err != nil {
			log.Println(err)
		}
		if hashPassword(oldPassword, salt) == password {
			return true
		}
	}
	return false
}

func updateEmail(userID int, email string) {
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
		if role == "student" {
			query = `UPDATE students SET email = ? WHERE student_id = ?`
			_, err = db.Query(query, email, userID)
			if err != nil {
				panic(err)
			}
		}
	}
}

func updatePassword(userID int, newPassword string) {
	var salt string
	query := `SELECT password_salt FROM login_details WHERE user_id = ?`
	lg, err := db.Query(query, userID)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&salt); err != nil {
			log.Println(err)
		}
		hashedPassword := hashPassword(newPassword, salt)
		query = `UPDATE login_details SET password = ? WHERE user_id = ?`
		_, err = db.Query(query, hashedPassword, userID)
	}
}

func clearUserSessions(userID int) {
	query := `DELETE FROM sessions WHERE user_id = ?`
	_, err := db.Query(query, userID)
	if err != nil {
		panic(err)
	}
}
