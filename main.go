package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var counter = 0

// increase the counter by one
func addCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	counter += 1
	json.NewEncoder(w).Encode(counter)
}

// decrease the counter by 1
func subCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	counter -= 1
	json.NewEncoder(w).Encode(counter)
}

// return the current counter value
func getCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(counter)
}

func main() {
	// init the mux router
	r := mux.NewRouter()

	// Create route handlers
	r.HandleFunc("/api/counter", getCounter).Methods("GET")
	r.HandleFunc("/api/counter", addCounter).Methods("PUT")
	r.HandleFunc("/api/counter", subCounter).Methods("DELETE")

	//run the server
	s := &http.Server{
		Addr:           ":8082",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
