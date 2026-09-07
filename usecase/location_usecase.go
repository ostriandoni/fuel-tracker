package usecase

import (
	"fuel-tracker/model"
	"fuel-tracker/repository"
)

type LocationUsecase interface {
	Create(name string) (*model.Location, error)
	GetAll() ([]model.Location, error)
	GetByID(id uint) (*model.Location, error)
	Update(id uint, name string) (*model.Location, error)
	Delete(id uint) error
}

type locationUsecase struct{ repo repository.LocationRepository }

func NewLocationUsecase(repo repository.LocationRepository) LocationUsecase {
	return &locationUsecase{repo}
}

func (u *locationUsecase) Create(name string) (*model.Location, error) {
	l := &model.Location{Name: name}
	if err := u.repo.Create(l); err != nil {
		return nil, err
	}
	return l, nil
}
func (u *locationUsecase) GetAll() ([]model.Location, error)        { return u.repo.FindAll() }
func (u *locationUsecase) GetByID(id uint) (*model.Location, error) { return u.repo.FindByID(id) }
func (u *locationUsecase) Update(id uint, name string) (*model.Location, error) {
	l, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	l.Name = name
	if err := u.repo.Update(l); err != nil {
		return nil, err
	}
	return l, nil
}
func (u *locationUsecase) Delete(id uint) error { return u.repo.Delete(id) }
