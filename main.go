package main

import (
	"log"
	"net/http"

	"github.com/Samuel200018/pills_backend/home"
	"github.com/gorilla/mux"
)

func main() {
	log.Print("Starting server")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home.HomeHandler)

	http.ListenAndServe(":8080", router)
}
