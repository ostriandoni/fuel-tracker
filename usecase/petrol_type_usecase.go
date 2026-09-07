package usecase

import (
	"fuel-tracker/model"
	"fuel-tracker/repository"

	"github.com/shopspring/decimal"
)

type PetrolTypeUsecase interface {
	Create(locationID uint, name string, price decimal.Decimal) (*model.PetrolType, error)
	GetAll() ([]model.PetrolType, error)
	GetByLocationID(locationID uint) ([]model.PetrolType, error)
	GetByID(id uint) (*model.PetrolType, error)
	Update(id uint, locationID uint, name string, price decimal.Decimal) (*model.PetrolType, error)
	Delete(id uint) error
}

type petrolTypeUsecase struct {
	repo repository.PetrolTypeRepository
}

func NewPetrolTypeUsecase(repo repository.PetrolTypeRepository) PetrolTypeUsecase {
	return &petrolTypeUsecase{repo}
}

func (u *petrolTypeUsecase) Create(locationID uint, name string, price decimal.Decimal) (*model.PetrolType, error) {
	p := &model.PetrolType{LocationID: locationID, Name: name, Price: price}
	if err := u.repo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}
func (u *petrolTypeUsecase) GetAll() ([]model.PetrolType, error) { return u.repo.FindAll() }
func (u *petrolTypeUsecase) GetByLocationID(locationID uint) ([]model.PetrolType, error) {
	return u.repo.FindByLocationID(locationID)
}
func (u *petrolTypeUsecase) GetByID(id uint) (*model.PetrolType, error) { return u.repo.FindByID(id) }
func (u *petrolTypeUsecase) Update(id uint, locationID uint, name string, price decimal.Decimal) (*model.PetrolType, error) {
	p, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	p.LocationID = locationID
	p.Name = name
	p.Price = price
	if err := u.repo.Update(p); err != nil {
		return nil, err
	}
	return p, nil
}
func (u *petrolTypeUsecase) Delete(id uint) error { return u.repo.Delete(id) }
