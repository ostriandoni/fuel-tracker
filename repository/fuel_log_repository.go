package repository

import (
	"fuel-tracker/model"

	"gorm.io/gorm"
)

type FuelLogRepository interface {
	Create(log *model.FuelLog) error
	FindAll() ([]model.FuelLog, error)
	FindByID(id uint) (*model.FuelLog, error)
	Update(log *model.FuelLog) error
	Delete(id uint) error
	Restore(id uint) error                        // new
	FindAllWithTrashed() ([]model.FuelLog, error) // new
	PermanentDelete(id uint) error                // new
}

type fuelLogRepository struct {
	db *gorm.DB
}

func NewFuelLogRepository(db *gorm.DB) FuelLogRepository {
	return &fuelLogRepository{db: db}
}

func (r *fuelLogRepository) Create(log *model.FuelLog) error {
	return r.db.Create(log).Error
}

func (r *fuelLogRepository) FindAll() ([]model.FuelLog, error) {
	var logs []model.FuelLog
	err := r.db.Order("date asc").Find(&logs).Error
	return logs, err
}

func (r *fuelLogRepository) FindByID(id uint) (*model.FuelLog, error) {
	var log model.FuelLog
	err := r.db.First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *fuelLogRepository) Update(log *model.FuelLog) error {
	return r.db.Save(log).Error
}

func (r *fuelLogRepository) Delete(id uint) error {
	return r.db.Delete(&model.FuelLog{}, id).Error
}

func (r *fuelLogRepository) Restore(id uint) error {
	return r.db.Unscoped().Model(&model.FuelLog{}).Where("id = ?", id).Update("deleted_at", nil).Error
}

func (r *fuelLogRepository) FindAllWithTrashed() ([]model.FuelLog, error) {
	var logs []model.FuelLog
	err := r.db.Unscoped().Order("date asc").Find(&logs).Error
	return logs, err
}

func (r *fuelLogRepository) PermanentDelete(id uint) error {
	return r.db.Unscoped().Delete(&model.FuelLog{}, id).Error
}
