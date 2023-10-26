package application

import "github.com/Samuel200018/pills_backend/medicine/domain"

type MedicineUseCase struct {
	medicineRepo domain.MedicineRepository
}

func NewUseCasesMedicine(medicineRepo domain.MedicineRepository) *MedicineUseCase {
	return &MedicineUseCase{medicineRepo}
}

func (muc *MedicineUseCase) CreateMedicine(medicine domain.Medicine) (domain.Medicine, error) {
	return muc.medicineRepo.Create(medicine)
}

func (muc *MedicineUseCase) GetMedicineById(id string) (domain.Medicine, error) {
	return muc.medicineRepo.Get(id)
}

func (muc *MedicineUseCase) GetMedicines(houseId uint) ([]domain.Medicine, error) {
	return muc.medicineRepo.GetAll(houseId)
}

func (muc *MedicineUseCase) UpdateMedicine(medicine domain.Medicine) (domain.Medicine, error) {
	return muc.medicineRepo.Update(medicine)
}

func (muc *MedicineUseCase) DeleteMedicine(medicine domain.Medicine) error {
	return muc.medicineRepo.Delete(medicine)
}
