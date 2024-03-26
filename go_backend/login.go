package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/argon2"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Response interface
type Response struct {
	LoginStatus bool   `json:"login_status"`
	SessionKey  string `json:"session_key"`
	ExpireTime  string `json:"expire_time"`
}

func ExecuteLogin(w http.ResponseWriter, r *http.Request) {
	// Get arguments from URL
	queryParams := r.URL.Query()
	login := queryParams.Get("login")
	password := queryParams.Get("password")
	fmt.Println("> User [ " + login + "] is trying to log in with password [ " + password + " ].")

	// Get salt from database
	salt := getUserSalt(login)

	// Hash password
	var hashedPassword = hashPassword(password, salt)

	// Check if login is correct and get session
	isLoggedIn, userId := tryLogin(login, hashedPassword, salt)

	// If login is correct, create session
	var sessionKey, expireTime string
	if isLoggedIn {
		sessionKey, expireTime = createSession(userId)
	}

	// Send response
	var response = Response{LoginStatus: isLoggedIn, SessionKey: sessionKey, ExpireTime: expireTime}
	json.NewEncoder(w).Encode(response)

	fmt.Println("> User [ " + login + "] logged in: [ " + strconv.FormatBool(isLoggedIn) + " ].")
}

func tryLogin(login, hashedPassword, salt string) (isLoggedIn bool, userId string) {
	// Get login details from database
	var query = `SELECT user_id, password FROM login_details WHERE login = ?`
	lg, err := db.Query(query, login)
	if err != nil {
		panic(err)
	}

	// Check if login is correct
	var databasePassword string
	for lg.Next() {
		if err = lg.Scan(&userId, &databasePassword); err != nil {
			log.Println(err)
		}

		if hashedPassword == databasePassword {
			isLoggedIn = true
		}
	}

	return isLoggedIn, userId
}

func createSession(userId string) (sessionKey, expireTime string) {
	// Generate session key - cryptographic sha256 hash
	for {
		sessionKey, _ = generateRandomString(256)
		if !sessionExists(sessionKey) {
			break
		}
	}

	var expirationTime = time.Now().Add(time.Hour * 24)
	var formattedExpirationTime = expirationTime.Format("2006-01-02 15:04:05")

	query := `INSERT INTO sessions(session_key, user_id, expire_time) VALUES (?, ?, ?)`
	_, err := db.ExecContext(context.Background(), query, sessionKey, userId, formattedExpirationTime)
	if err != nil {
		panic(err)
	}

	return sessionKey, formattedExpirationTime
}

func sessionExists(newSessionKey string) bool {
	// Check if session exists
	var query = `SELECT session_key FROM sessions WHERE session_key = ?`
	lg, err := db.Query(query, newSessionKey)
	if err != nil {
		panic(err)
	}

	// Check if session exists
	var sessionKey = ""
	for lg.Next() {
		if err = lg.Scan(&sessionKey); err != nil {
			log.Println(err)
		}

		if sessionKey == newSessionKey {
			return true
		}
	}
	return false
}

func getUserSalt(login string) string {
	var passwordSalt = ""

	var query = `SELECT password_salt FROM login_details WHERE login = ?`
	lg, err := db.Query(query, login)
	if err != nil {
		panic(err)
	}

	for lg.Next() {
		if err = lg.Scan(&passwordSalt); err != nil {
			log.Println(err)
		}
	}
	return passwordSalt
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func hashPassword(password string, salt string) string {
	key := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	return hex.EncodeToString(key)
}
