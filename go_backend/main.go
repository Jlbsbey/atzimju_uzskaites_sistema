package main

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/argon2"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

func main() {
	cfg := mysql.NewConfig()
	(*cfg).User = "root"
	(*cfg).Addr = "localhost"
	(*cfg).Passwd = "root"
	(*cfg).Net = "tcp"
	(*cfg).DBName = "grade"
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/login", ExecLogin).Methods("GET")
	router.HandleFunc("/home", HomePage).Methods("GET")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func returnSomething(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode("Hello World !, user " + strconv.Itoa(id))
}

func hashPassword(password string, salt string) string {
	key := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	return hex.EncodeToString(key)
}
