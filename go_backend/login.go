package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ExecLogin(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	URLlog := queryParams.Get("login")
	URLpas := queryParams.Get("password")
	salt := getSalt(URLlog)
	hashedPasw := hashPassword(URLpas, salt)
	session := checkLogin(hashedPasw, URLlog, salt)
	if session == "" {
		responseString := "auth:-1"
		w.Write([]byte(responseString))
		return
	}
	fmt.Println(session)
	responseString := fmt.Sprintf("auth: %s", session)
	w.Write([]byte(responseString))
}

func checkLogin(passw string, login string, salt string) string {
	DBlogin := ""
	DBpassw := ""
	query := `SELECT studID, password FROM student WHERE studID = ? AND password = ?`
	lg, err := db.Query(query, login, passw)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&DBlogin, &DBpassw); err != nil {
			log.Println(err)
		}
		if login == DBlogin && passw == DBpassw {
			currentTime := time.Now()
			sessionNum := hashPassword(currentTime.String(), salt)
			sessiontoDB(sessionNum, currentTime, login)
			return sessionNum
		}
	}
	if DBlogin == "" {
		query := `SELECT studID, password FROM prof WHERE profID = ? AND password = ?`
		lg, err := db.Query(query, login)
		if err != nil {
			panic(err)
		}
		for lg.Next() {
			if err = lg.Scan(&DBlogin, &DBpassw); err != nil {
				log.Println(err)
			}
			if login == DBlogin && passw == DBpassw {
				if err = lg.Scan(&DBlogin, &DBpassw); err != nil {
					log.Println(err)
				}
				if login == DBlogin && passw == DBpassw {
					currentTime := time.Now()
					sessionNum := hashPassword(currentTime.String(), salt)
					sessiontoDB(sessionNum, currentTime, login)
					return sessionNum
				}
			}
		}
	}
	return ""
}

func sessiontoDB(session string, now time.Time, login string) {
	query := `INSERT INTO sessions(sessionID, expire, login) VALUES (?, ?)`
	expirationTime := now.Add(1 * time.Hour)
	formatted := expirationTime.Format("2006-01-02 15:04:05")
	_, err := db.ExecContext(context.Background(), query, session, formatted, login)
	if err != nil {
		panic(err)
	}
}

func getSalt(login string) string {
	salt := ""
	query := `SELECT salt FROM student WHERE studID = ?`
	lg, err := db.Query(query, login)
	if err != nil {
		panic(err)
	}
	for lg.Next() {
		if err = lg.Scan(&salt); err != nil {
			log.Println(err)
		}
	}
	if salt == "" {
		query = `SELECT salt FROM profs WHERE profID = ?`
		lg, err = db.Query(query, login)
		if err != nil {
			panic(err)
		}
		for lg.Next() {
			if err = lg.Scan(&salt); err != nil {
				log.Println(err)
			}
		}
	}
	return salt
}
