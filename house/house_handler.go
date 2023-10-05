package house

import (
	"encoding/json"
	authApp "github.com/Samuel200018/pills_backend/auth/application"
	"github.com/Samuel200018/pills_backend/house/application"
	"github.com/Samuel200018/pills_backend/house/domain"
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
		"message":      "House created successfully",
		"user_updated": userUpdated,
	}
	json.NewEncoder(w).Encode(response)
}
