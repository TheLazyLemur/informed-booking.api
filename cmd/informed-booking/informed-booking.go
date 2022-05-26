package main

import (
	"fmt"
	"log"
	"net/http"

	booking "github.com/TheLazyLemur/informed-booking.api/internal/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

    booking.Router(router)

	fmt.Println("Serving informed-booking.api on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
