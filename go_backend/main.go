package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB

type Response_Body struct {
	Status   string `json:"status"`
	Error    string `json:"error"`
	Response any    `json:"content"`
}

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
	go ClearSessions()

	// Create a new router
	router := mux.NewRouter()

	// Add CORS middleware
	router.Use(corsMiddleware)

	router.HandleFunc("/login", ExecuteLogin).Methods("GET")
	router.HandleFunc("/home", HomePage).Methods("GET")
	router.HandleFunc("/profile", ProfilePage).Methods("GET")
	router.HandleFunc("/changeData", ChangeData).Methods("GET")
	//router.HandleFunc("/studentList", StudentList).Methods("GET")
	router.HandleFunc("/addMark", AddMark).Methods("GET")
	router.HandleFunc("/addSubject", AddSubject).Methods("GET")
	//router.HandleFunc("/changeUserData", ChangeUserData).Methods("GET")
	router.HandleFunc("/changeUserSubjects", ChangeUserSubjects).Methods("GET")
	router.HandleFunc("/searchSubjects", SearchSubjects).Methods("GET")
	router.HandleFunc("/addUser", AddUser).Methods("GET")

	var development = false

	if development {
		// Start the HTTP server
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		fmt.Println("Starting server in working mode...")
		// Start the HTTP server
		log.Fatal(http.ListenAndServeTLS(
			":8443",
			"fullchain.crt",
			"privkey.key",
			router,
		))
	}
}

// Define the CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers to allow cross-origin requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// If it's a preflight request, send a 200 OK status
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
