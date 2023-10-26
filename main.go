package main

import (
	"github.com/Samuel200018/pills_backend/auth"
	"github.com/Samuel200018/pills_backend/auth/application"
	"github.com/Samuel200018/pills_backend/auth/domain"
	"github.com/Samuel200018/pills_backend/auth/infrastructure"
	"github.com/Samuel200018/pills_backend/db"
	"github.com/Samuel200018/pills_backend/house"
	appHouse "github.com/Samuel200018/pills_backend/house/application"
	domHouse "github.com/Samuel200018/pills_backend/house/domain"
	infraHouse "github.com/Samuel200018/pills_backend/house/infrastructure"
	"github.com/Samuel200018/pills_backend/medicine"
	appMedicine "github.com/Samuel200018/pills_backend/medicine/application"
	domMedicine "github.com/Samuel200018/pills_backend/medicine/domain"
	infaMedicine "github.com/Samuel200018/pills_backend/medicine/infrastructure"
	"log"
	"net/http"

	"github.com/Samuel200018/pills_backend/home"
	"github.com/gorilla/mux"
)

func main() {
	log.Print("Starting server")
	db.Connection()

	err := db.DB.AutoMigrate(&domain.User{})
	err = db.DB.AutoMigrate(&domHouse.House{})
	err = db.DB.AutoMigrate(&domMedicine.Medicine{})

	if err != nil {
		log.Print("Error auto migrate")
		return
	}

	router := mux.NewRouter().StrictSlash(true)

	userRepository := infrastructure.NewDatabaseAuthRepository(db.DB)
	houseRepository := infraHouse.NewDatabaseHouseRepository(db.DB)
	medicineRepository := infaMedicine.NewDatabaseMedicineRepository(db.DB)

	userUseCases := application.NewUseCasesAuth(userRepository)
	houseUseCases := appHouse.NewUseCasesHouse(houseRepository)
	medicineUseCases := appMedicine.NewUseCasesMedicine(medicineRepository)

	authHandler := auth.NewAuthHandler(*userUseCases)
	houseHandler := house.NewHouseHandler(*houseUseCases, *userUseCases)
	medicineHandler := medicine.NewMedicineHandler(*medicineUseCases, *houseUseCases)

	router.HandleFunc("/", home.HomeHandler)

	//Auth
	router.HandleFunc("/createUser", authHandler.CreateUser).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("GET")

	//Protected routes
	//Auth
	protectedRoutes := router.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(auth.JwtVerify)
	protectedRoutes.HandleFunc("/check-status", authHandler.CheckStatus).Methods("GET")
	protectedRoutes.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	//House
	protectedRoutes.HandleFunc("/create-house", houseHandler.CreateHouse).Methods("POST")
	protectedRoutes.HandleFunc("/join-house/{house-id}", houseHandler.JoinHouse).Methods("PUT")
	protectedRoutes.HandleFunc("/exit-house/{house-id}", houseHandler.ExitHouse).Methods("DELETE")
	//Medicine
	protectedRoutes.HandleFunc("/create-medicine", medicineHandler.CreateMedicine).Methods("POST")
	protectedRoutes.HandleFunc("/get-all-medicines/{house-id}", medicineHandler.GetAllMedicinesByHouseId).Methods("GET")
	protectedRoutes.HandleFunc("/update-medicine", medicineHandler.UpdateMedicine).Methods("PUT")

	log.Fatal(
		http.ListenAndServe(":3800", router),
	)

}
