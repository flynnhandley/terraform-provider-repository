package main

import (
	"log"
	"net/http"

	"./src/tpr"
	"github.com/gorilla/mux"
)

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/provider/{id}", tpr.GetProvider).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
