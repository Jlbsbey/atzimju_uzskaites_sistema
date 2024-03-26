package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

func main() {
	cfg := mysql.NewConfig()
	(*cfg).User = "remote"
	(*cfg).Addr = "104.248.86.80:3306"
	(*cfg).Passwd = "ts32YQ?!Twa2$Ej"
	(*cfg).Net = "tcp"
	(*cfg).DBName = "grade"
	(*cfg).ParseTime = true
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	/*session := ExecLogin()
	print(session)
	HomePage(session)*/
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
