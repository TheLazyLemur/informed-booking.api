package booking

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	router.HandleFunc("/", get).Methods("GET")
	router.HandleFunc("/", post).Methods("POST")
	router.HandleFunc("/", put).Methods("PUT")
	router.HandleFunc("/", delete).Methods("DELETE")
}

func get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello,World")
}

func post(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello,World")
}

func put(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello,World")
}

func delete(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello,World")
}
