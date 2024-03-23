package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define the endpoints for CRUD operations
	router.HandleFunc("/items/{id}", returnSomething).Methods("GET")

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
