package medicine

import (
	"encoding/json"
	appHouse "github.com/Samuel200018/pills_backend/house/application"
	"github.com/Samuel200018/pills_backend/medicine/application"
	"github.com/Samuel200018/pills_backend/medicine/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type MedHandler struct {
	medicineUseCases application.MedicineUseCase
	houseUseCases    appHouse.HouseUseCase
}

func NewMedicineHandler(medUseCases application.MedicineUseCase, houseUseCases appHouse.HouseUseCase) *MedHandler {
	return &MedHandler{medUseCases, houseUseCases}
}

func (medHandler *MedHandler) CreateMedicine(w http.ResponseWriter, r *http.Request) {

	dose, err := strconv.ParseFloat(r.FormValue("dose"), 64)

	if err != nil {
		http.Error(w, "Error parsing dose", http.StatusBadRequest)
		return
	}

	amountAvailable, err := strconv.ParseFloat(r.FormValue("amountAvailable"), 64)

	if err != nil {
		http.Error(w, "Error parsing amount available", http.StatusBadRequest)
		return
	}

	dueDate, err := time.Parse("2006-01-02", r.FormValue("dueDate"))

	if err != nil {
		http.Error(w, "Error parsing due date", http.StatusBadRequest)
		return
	}

	house, err := medHandler.houseUseCases.GetHouseById(r.FormValue("houseId"))
	if err != nil {
		http.Error(w, "Error house not found", http.StatusNotFound)
		return
	}

	medicine := domain.Medicine{
		Name:            r.FormValue("name"),
		Dose:            dose,
		AmountAvailable: amountAvailable,
		DueDate:         dueDate,
		ActiveCompounds: r.FormValue("activeCompounds"),
		HouseID:         house.ID,
	}

	createMedicine, err := medHandler.medicineUseCases.CreateMedicine(medicine)
	if err != nil {
		http.Error(w, "Error creating medicine", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"message":  "Medicine created successfully",
		"medicine": createMedicine,
	}

	json.NewEncoder(w).Encode(response)
}

func (medHandler *MedHandler) GetAllMedicinesByHouseId(w http.ResponseWriter, r *http.Request) {
	houseFromRequest := mux.Vars(r)["house_id"]

	house, err := medHandler.houseUseCases.GetHouseById(houseFromRequest)
	if err != nil {
		http.Error(w, "House not found", http.StatusNotFound)
		return
	}

	medicines, err := medHandler.medicineUseCases.GetMedicines(house.ID)
	if err != nil {
		return
	}
	response := map[string]interface{}{
		"message":   "Get all medicines successfully",
		"medicines": medicines,
	}

	json.NewEncoder(w).Encode(response)
}

func (medHandler *MedHandler) UpdateMedicine(w http.ResponseWriter, r *http.Request) {

}
