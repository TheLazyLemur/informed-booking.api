package main

import (
	"github.com/gorilla/mux"
    "github.com/TheLazyLemur/informed-booking.api/internal"
    "fmt"
    "log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

    router.HandleFunc("/", greeting.Hello).Methods("GET")

    fmt.Println("Serving informed-booking.api on 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
