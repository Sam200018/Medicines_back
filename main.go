package main

import (
	"github.com/Samuel200018/pills_backend/auth"
	"github.com/Samuel200018/pills_backend/auth/application"
	"github.com/Samuel200018/pills_backend/auth/domain"
	"github.com/Samuel200018/pills_backend/auth/infrastructure"
	"github.com/Samuel200018/pills_backend/db"
	"log"
	"net/http"

	"github.com/Samuel200018/pills_backend/home"
	"github.com/gorilla/mux"
)

func main() {
	log.Print("Starting server")
	db.Connection()

	err := db.DB.AutoMigrate(&domain.User{})
	if err != nil {
		log.Print("Error auto migrate")
		return
	}

	router := mux.NewRouter().StrictSlash(true)

	userRepository := infrastructure.NewDatabaseAuthRepository(db.DB)

	userUseCases := application.NewUseCasesAuth(userRepository)

	authHandler := auth.NewAuthHandler(*userUseCases)

	router.HandleFunc("/", home.HomeHandler)

	//Auth
	router.HandleFunc("/createUser", authHandler.CreateUser).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("GET")

	//Protected routes
	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(auth.JwtVerify)
	protectedRoutes.HandleFunc("/check-status", authHandler.CheckStatus).Methods("GET")
	protectedRoutes.HandleFunc("/logout", authHandler.Logout).Methods("POST")

	log.Fatal(
		http.ListenAndServe(":3800", router),
	)

}
