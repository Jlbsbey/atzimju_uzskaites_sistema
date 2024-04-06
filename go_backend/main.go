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

	var development = true

	if development {
		// Start the HTTP server
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		// Start the HTTP server
		log.Fatal(http.ListenAndServeTLS(
			":8443",
			"fullchain.crt",
			"privkey.key",
			router,
		))
	}
}

//6f87d01d35eb5aace608c08632745f3bed9b32613df6c777150c04cef3e86ba0c2ac3f6e74edf5e286ba6e353f78b99a7b27285904d22d6a17391fd3ebd5e4a8ee8f86ab984d4d704ec9d2e2a3c38180edee0ddd32780e430f52f542c5c0ed05cd6339753a59bf1406d4e3b339d2e403b3a5806f687cf675af932a73549ac63c

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
