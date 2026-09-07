package repository

import (
	"fuel-tracker/model"

	"gorm.io/gorm"
)

type LocationRepository interface {
	Create(l *model.Location) error
	FindAll() ([]model.Location, error)
	FindByID(id uint) (*model.Location, error)
	Update(l *model.Location) error
	Delete(id uint) error
}

type locationRepository struct{ db *gorm.DB }

func NewLocationRepository(db *gorm.DB) LocationRepository { return &locationRepository{db} }

func (r *locationRepository) Create(l *model.Location) error { return r.db.Create(l).Error }
func (r *locationRepository) FindAll() ([]model.Location, error) {
	var locs []model.Location
	err := r.db.Order("name asc").Find(&locs).Error
	return locs, err
}
func (r *locationRepository) FindByID(id uint) (*model.Location, error) {
	var l model.Location
	err := r.db.First(&l, id).Error
	if err != nil {
		return nil, err
	}
	return &l, nil
}
func (r *locationRepository) Update(l *model.Location) error { return r.db.Save(l).Error }
func (r *locationRepository) Delete(id uint) error           { return r.db.Delete(&model.Location{}, id).Error }
