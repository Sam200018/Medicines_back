package house

import (
	"encoding/json"
	authApp "github.com/Samuel200018/pills_backend/auth/application"
	"github.com/Samuel200018/pills_backend/house/application"
	"github.com/Samuel200018/pills_backend/house/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		"message":      "House created successfully",
		"user_updated": userUpdated,
	}
	json.NewEncoder(w).Encode(response)
}

func (houseHad *HouseHandler) JoinHouse(w http.ResponseWriter, r *http.Request) {
	houseFromRequest := mux.Vars(r)["house-id"]
	parseUint, err := strconv.ParseUint(houseFromRequest, 10, 32)
	if err != nil {
		http.Error(w, "Error parsing string", http.StatusBadRequest)
		return
	}
	houseId := uint(parseUint)
	userId := r.FormValue("user_id")

	userUpdated, err := houseHad.authUsesCases.AddHouse(userId, houseId)
	if err != nil {
		http.Error(w, "Error joining to house", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message":      "House created successfully",
		"user_updated": userUpdated,
	}
	json.NewEncoder(w).Encode(response)

}

func (houseHad *HouseHandler) ExitHouse(w http.ResponseWriter, r *http.Request) {
	houseFromRequest := mux.Vars(r)["house-id"]
	parseUint, err := strconv.ParseUint(houseFromRequest, 10, 32)
	if err != nil {
		http.Error(w, "Error parsing string", http.StatusBadRequest)
		return
	}
	houseId := uint(parseUint)
	userId := r.FormValue("user_id")

	countUsers, err := houseHad.authUsesCases.UsersWithHouseId(houseId)
	if err != nil {
		http.Error(w, "Error counting users with house id", http.StatusBadRequest)
		return
	}

	if countUsers > 1 {
		//	Eliminar house\
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
		"message":      "Exit house successfully",
		"user_updated": userUpdated,
	}
	json.NewEncoder(w).Encode(response)

}
