package house

import (
	"encoding/json"
	authApp "github.com/Samuel200018/pills_backend/auth/application"
	"github.com/Samuel200018/pills_backend/house/application"
	"github.com/Samuel200018/pills_backend/house/domain"
	"github.com/gorilla/mux"
	"net/http"
)

type HouseHandler struct {
	houseUseCases application.HouseUseCase
	authUsesCases authApp.AuthUseCase
}

func NewHouseHandler(houseUseCase application.HouseUseCase, authUseCase authApp.AuthUseCase) *HouseHandler {
	return &HouseHandler{houseUseCase, authUseCase}
}

func (houseHad *HouseHandler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	house := domain.House{
		Name: r.FormValue("name"),
	}

	createHouse, err := houseHad.houseUseCases.CreateHouse(house)
	if err != nil {
		http.Error(w, "House not created", http.StatusFailedDependency)
		return
	}

	userId := r.FormValue("user_id")

	userUpdated, err := houseHad.authUsesCases.AddHouse(userId, createHouse.ID)

	if err != nil {
		http.Error(w, "User not updated", http.StatusConflict)
		return
	}

	response := map[string]interface{}{
		"message": "House created successfully",
		"user":    userUpdated,
	}
	json.NewEncoder(w).Encode(response)
}

func (houseHad *HouseHandler) JoinHouse(w http.ResponseWriter, r *http.Request) {
	houseFromRequest := mux.Vars(r)["house-id"]

	userId := r.FormValue("user_id")
	if userId == "" {
		http.Error(w, "Empty user id", http.StatusBadRequest)
		return
	}

	house, err := houseHad.houseUseCases.GetHouseById(houseFromRequest)
	if err != nil {
		http.Error(w, "House not found", http.StatusNotFound)
		return
	}

	userUpdated, err := houseHad.authUsesCases.AddHouse(userId, house.ID)
	if err != nil {
		http.Error(w, "Error joining to house", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "Joined to house successfully",
		"user":    userUpdated,
	}
	json.NewEncoder(w).Encode(response)

}

func (houseHad *HouseHandler) ExitHouse(w http.ResponseWriter, r *http.Request) {
	houseFromRequest := mux.Vars(r)["house-id"]

	userId := r.FormValue("user_id")
	if userId == "" {
		http.Error(w, "Empty user id", http.StatusBadRequest)
		return
	}

	house, err := houseHad.houseUseCases.GetHouseById(houseFromRequest)
	if err != nil {
		http.Error(w, "House not found", http.StatusNotFound)
		return
	}

	countUsers, err := houseHad.authUsesCases.UsersWithHouseId(house.ID)
	if err != nil {
		http.Error(w, "Error counting users with house id", http.StatusBadRequest)
		return
	}

	if countUsers == 1 {
		err := houseHad.houseUseCases.DeleteHouse(houseFromRequest)
		if err != nil {
			http.Error(w, "Error deleting house", http.StatusBadRequest)
			return
		}
	}

	userUpdated, err := houseHad.authUsesCases.ExitHouse(userId)

	if err != nil {
		print(err)
		http.Error(w, "User not exit the house", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "Exit house successfully",
		"user":    userUpdated,
	}
	json.NewEncoder(w).Encode(response)

}
