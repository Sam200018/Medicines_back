package domain

type MedicineRepository interface {
	Create(medicine Medicine) (Medicine, error)
	Get(id string) (Medicine, error)
	Update(medicine Medicine) (Medicine, error)
	Delete(medicine Medicine) error
	GetAll(houseId uint) ([]Medicine, error)
}
