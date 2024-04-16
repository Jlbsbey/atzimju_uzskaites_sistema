package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ChangeData(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")            //all
	email := queryParams.Get("email")             //all
	oldPassword := queryParams.Get("oldPassword") //all
	newPassword := queryParams.Get("newPassword") //all
	newName := queryParams.Get("newName")         //admin
	newSurname := queryParams.Get("newSurname")   //admin
	username := queryParams.Get("username")       //all
	userID := UserIDbyUsername(username)
	originUserID := getUserID(session)
	isAdmin := checkAdmin(originUserID)
	json.NewEncoder(w).Encode(isAdmin)
	sameUsers := userID == originUserID
	if originUserID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if email != "" && sameUsers {
		fmt.Println("1")
		updateData(originUserID, email, "email")
	} else if email != "" && isAdmin {
		updateData(userID, email, "email")
	}
	if newName != "" && sameUsers && isAdmin {
		updateData(originUserID, newName, "name")
	} else if newName != "" && isAdmin {
		updateData(userID, newName, "name")
	}
	if newSurname != "" && sameUsers && isAdmin {
		updateData(originUserID, newSurname, "surname")
	} else if newSurname != "" && isAdmin {
		updateData(userID, newSurname, "surname")
	}
	if oldPassword != "" && checkValidity(userID, oldPassword) && sameUsers {
		updatePassword(originUserID, newPassword)
		clearUserSessions(originUserID)
	} else if newPassword != "" && checkAdmin(userID) {
		updatePassword(userID, newPassword)
		clearUserSessions(userID)
	} else if newPassword != "" {
		var response = Response_Body{Status: "error", Error: "Password is incorrect"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	var response = Response_Body{Status: "OK", Error: ""} //истекло время сессии или пользователь не был найден по сессии
	json.NewEncoder(w).Encode(response)
	return

}

func UserIDbyUsername(username string) int {
	var userID int
	query := `SELECT user_id FROM login_details WHERE username = ?`
	lg, err := db.Query(query, username)
	for lg.Next() {
		if err = lg.Scan(&userID); err != nil {
			log.Println(err)
		}
	}
	return userID
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

func updateData(userID int, email string, dataType string) {
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
			query = fmt.Sprintf("UPDATE students SET %s = ? WHERE student_id = ?", dataType)
			_, err = db.Query(query, email, userID)
			if err != nil {
				panic(err)
			}
		} else if role == "professor" {
			query = fmt.Sprintf("UPDATE professors SET %s = ? WHERE professor_id = ?", dataType)
			_, err = db.Query(query, email, userID)
			if err != nil {
				panic(err)
			}
		} else if role == "admin" {
			query = `UPDATE configuration SET value = ? WHERE name = 'AdminEmail'`
			query = fmt.Sprintf("UPDATE configuration SET value = ? WHERE name = %s", "Admin"+strings.ToUpper(dataType[0:1])+dataType[1:])
			_, err = db.Query(query, email)
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
