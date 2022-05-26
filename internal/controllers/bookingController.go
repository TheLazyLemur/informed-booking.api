package booking

import (
	"fmt"
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
	fmt.Println("Get not implemented")
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post not implemented")
}

func put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Put not implemented")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete not implemented")
}
