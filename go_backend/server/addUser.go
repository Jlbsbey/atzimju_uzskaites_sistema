package server

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	session := queryParams.Get("auth")
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	role := queryParams.Get("role")
	email := strings.ToLower(queryParams.Get("email"))
	password := queryParams.Get("password")
	avatarLink := queryParams.Get("avatar_link")
	userID := getUserID(session)
	fmt.Println(name, " ", surname)
	if userID == -1 {
		var response = Response_Body{Status: "error", Error: "Session expired"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	if !checkAdmin(userID) {
		var response = Response_Body{Status: "error", Error: "User does not have sufficient rights"}
		json.NewEncoder(w).Encode(response)
		return
	}
	salt := hashPassword(time.Now().Format("2006-01-02 15:04:05"), "a")
	hashedPassword := hashPassword(password, salt)
	username := generateUsername(strings.ToLower(name), strings.ToLower(surname))
	if username == "" {
		var response = Response_Body{Status: "error", Error: "Wtf is this error"} //истекло время сессии или пользователь не был найден по сессии
		json.NewEncoder(w).Encode(response)
		return
	}
	createUser(name, surname, role, email, hashedPassword, salt, username, avatarLink)
	response := Response_Body{Status: "OK"}
	json.NewEncoder(w).Encode(response)
}

func createUser(name string, surname string, role string, email string, hashedPassword string, salt string, username string, avatarLink string) {
	userID := generateRandomInteger(1, math.MaxUint32)
	if role == "student" {
		query := `INSERT INTO students(student_id, name, surname, email, avatar_link) VALUES (?, ?, ?, ?, ?)`
		_, err := db.ExecContext(context.Background(), query, userID, name, surname, email, avatarLink)
		if err != nil {
			panic(err)
		}
	} else if role == "professor" {
		query := `INSERT INTO professors(professor_id, name, surname, email, avatar_link) VALUES (?, ?, ?, ?, ?)`
		_, err := db.ExecContext(context.Background(), query, userID, name, surname, email, avatarLink)
		if err != nil {
			panic(err)
		}
	}
	query := `INSERT INTO login_details(user_id, role, username, password, password_salt) VALUES (?, ?, ?, ?, ?)`
	_, err := db.ExecContext(context.Background(), query, userID, role, username, hashedPassword, salt)
	if err != nil {
		panic(err)
	}
}

func generateUsername(name string, surname string) string {
	counter := 0
	query := `SELECT name FROM students WHERE name = ? AND surname = ?`
	lg, err := db.Query(query, name, surname)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		counter++
	}
	query = `SELECT name FROM professors WHERE name = ? AND surname = ?`
	lg, err = db.Query(query, name, surname)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		counter++
	}
	if counter > 1 && len(surname) >= 10 {
		return name[0:1] + "." + surname[:10] + fmt.Sprint(counter)
	} else if counter == 0 && len(surname) >= 10 {
		return name[0:1] + "." + surname[:10]
	} else if counter > 0 && len(surname) < 10 {
		return name[0:1] + "." + surname[:] + fmt.Sprint(counter)
	} else if counter == 0 && len(surname) < 10 {
		return name[0:1] + "." + surname[:]
	}
	return ""
}
