package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ChangeData(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	email := queryParams.Get("email")
	oldPassword := queryParams.Get("oldPassword")
	newPassword := queryParams.Get("newPassword")
	newName := queryParams.Get("newName")
	newSurname := queryParams.Get("newSurname")
	userID := getUserID(session)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if oldPassword == "" {
		if email != "" {
			updateEmail(userID, email)
		}
		if newName != "" && checkAdmin(userID) {
			UpdateName(userID, newName)
		}
		if newSurname != "" && checkAdmin(userID) {
			UpdateSurname(userID, newName)
		}
		var response = Response_Body{Status: "OK", Error: ""} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if checkValidity(userID, oldPassword) {
		if email != "" {
			updateEmail(userID, email)
		}
		if newName != "" && checkAdmin(userID) {
			UpdateName(userID, newName)
		}
		if newSurname != "" && checkAdmin(userID) {
			UpdateSurname(userID, newName)
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

func UpdateSurname(userID int, surname string) {
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
			query = `UPDATE students SET surname = ? WHERE student_id = ?`
			_, err = db.Query(query, surname, userID)
			if err != nil {
				panic(err)
			}
		} else if role == "professors" {
			query = `UPDATE professors SET surname = ? WHERE professor_id = ?`
			_, err = db.Query(query, surname, userID)
			if err != nil {
				panic(err)
			}
		} else if role == "admin" {
			query = `UPDATE admins SET surname = ? WHERE admin_id = ?`
			_, err = db.Query(query, surname, userID)
			if err != nil {
				panic(err)
			}
		}
	}
}

func UpdateName(userID int, name string) {
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
			query = `UPDATE students SET name = ? WHERE student_id = ?`
			_, err = db.Query(query, name, userID)
			if err != nil {
				panic(err)
			}
		} else if role == "professors" {
			query = `UPDATE professors SET name = ? WHERE professor_id = ?`
			_, err = db.Query(query, name, userID)
			if err != nil {
				panic(err)
			}
		} else if role == "admin" {
			query = `UPDATE admins SET name = ? WHERE admin_id = ?`
			_, err = db.Query(query, name, userID)
			if err != nil {
				panic(err)
			}
		}
	}
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
		} else if role == "professors" {
			query = `UPDATE professors SET email = ? WHERE professor_id = ?`
			_, err = db.Query(query, email, userID)
			if err != nil {
				panic(err)
			}
		} else if role == "admin" {
			query = `UPDATE admins SET email = ? WHERE admin_id = ?`
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
