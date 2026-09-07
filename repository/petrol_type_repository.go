package repository

import (
	"fuel-tracker/model"

	"gorm.io/gorm"
)

type PetrolTypeRepository interface {
	Create(p *model.PetrolType) error
	FindAll() ([]model.PetrolType, error)
	FindByLocationID(locationID uint) ([]model.PetrolType, error)
	FindByID(id uint) (*model.PetrolType, error)
	Update(p *model.PetrolType) error
	Delete(id uint) error
}

type petrolTypeRepository struct{ db *gorm.DB }

func NewPetrolTypeRepository(db *gorm.DB) PetrolTypeRepository { return &petrolTypeRepository{db} }

func (r *petrolTypeRepository) Create(p *model.PetrolType) error { return r.db.Create(p).Error }
func (r *petrolTypeRepository) FindAll() ([]model.PetrolType, error) {
	var p []model.PetrolType
	err := r.db.Preload("Location").Find(&p).Error
	return p, err
}
func (r *petrolTypeRepository) FindByLocationID(locationID uint) ([]model.PetrolType, error) {
	var p []model.PetrolType
	err := r.db.Where("location_id = ?", locationID).Order("name asc").Find(&p).Error
	return p, err
}
func (r *petrolTypeRepository) FindByID(id uint) (*model.PetrolType, error) {
	var p model.PetrolType
	err := r.db.Preload("Location").First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *petrolTypeRepository) Update(p *model.PetrolType) error { return r.db.Save(p).Error }
func (r *petrolTypeRepository) Delete(id uint) error {
	return r.db.Delete(&model.PetrolType{}, id).Error
}
