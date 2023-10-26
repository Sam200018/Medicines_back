package application

import "github.com/Samuel200018/pills_backend/house/domain"

type HouseUseCase struct {
	homeRepo domain.HouseRepository
}

func NewUseCasesHouse(homeRepo domain.HouseRepository) *HouseUseCase {
	return &HouseUseCase{homeRepo}
}

func (huc *HouseUseCase) CreateHouse(house domain.House) (domain.House, error) {
	return huc.homeRepo.Create(house)
}

func (huc *HouseUseCase) GetHouseById(id string) (domain.House, error) {
	return huc.homeRepo.GetById(id)
}

func (huc *HouseUseCase) UpdateHouse(house domain.House) (domain.House, error) {
	return huc.homeRepo.Update(house)
}

func (huc *HouseUseCase) DeleteHouse(id string) error {
	return huc.homeRepo.Delete(id)
}
