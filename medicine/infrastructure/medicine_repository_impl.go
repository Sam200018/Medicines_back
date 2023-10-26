package infrastructure

import (
	"github.com/Samuel200018/pills_backend/medicine/domain"
	"gorm.io/gorm"
)

type DatabaseMedicineRepository struct {
	db *gorm.DB
}

func NewDatabaseMedicineRepository(db *gorm.DB) domain.MedicineRepository {
	return &DatabaseMedicineRepository{
		db,
	}
}

func (d *DatabaseMedicineRepository) GetAll(houseId uint) ([]domain.Medicine, error) {
	var medicines []domain.Medicine

	err := d.db.Find(&medicines).Where("houseId =?", houseId).Error
	return medicines, err
}

func (d *DatabaseMedicineRepository) Create(medicine domain.Medicine) (domain.Medicine, error) {
	err := d.db.Create(&medicine).Error
	return medicine, err
}

func (d *DatabaseMedicineRepository) Get(id string) (domain.Medicine, error) {
	var medicine domain.Medicine
	err := d.db.First(&medicine, id).Error

	return medicine, err
}

func (d *DatabaseMedicineRepository) Update(medicine domain.Medicine) (domain.Medicine, error) {
	err := d.db.Model(&domain.Medicine{}).Where("id =?", medicine.ID).Updates(&medicine).Error
	return medicine, err
}

func (d *DatabaseMedicineRepository) Delete(medicine domain.Medicine) error {
	d.db.First(&medicine)

	return d.db.Delete(&medicine).Error
}
