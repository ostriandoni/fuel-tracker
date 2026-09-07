package usecase

import (
	"fuel-tracker/dto"
	"fuel-tracker/helper"
	"fuel-tracker/model"
	"fuel-tracker/repository"
)

type FuelLogUsecase interface {
	Create(req dto.CreateFuelLogRequest) (*dto.FuelLogResponse, error)
	GetAll() ([]dto.FuelLogResponse, error)
	GetByID(id uint) (*dto.FuelLogResponse, error)
	Update(id uint, req dto.UpdateFuelLogRequest) (*dto.FuelLogResponse, error)
	Delete(id uint) error
}

type fuelLogUsecase struct {
	repo repository.FuelLogRepository
}

func NewFuelLogUsecase(repo repository.FuelLogRepository) FuelLogUsecase {
	return &fuelLogUsecase{repo: repo}
}

func (u *fuelLogUsecase) toResponse(m *model.FuelLog) *dto.FuelLogResponse {
	distance := helper.CalculateDistance(m.KmStart, m.KmEnd)
	paidPerKm := helper.CalculatePaidPerKm(m.TotalPaid, distance)
	usageKmL := helper.CalculateUsageKmL(distance, m.LitersFilled)

	return &dto.FuelLogResponse{
		ID:             m.ID,
		Month:          int(m.Date.Month()),
		Date:           m.Date,
		PricePerLiter:  m.PricePerLiter,
		TotalPaid:      m.TotalPaid,
		LitersFilled:   m.LitersFilled,
		KmStart:        m.KmStart,
		KmEnd:          m.KmEnd,
		Distance:       distance,
		PaidPerKm:      paidPerKm,
		UsageKmL:       usageKmL,
		LocationID:     m.LocationID,
		LocationName:   m.Location.Name, // struct field, not the struct itself
		PetrolTypeID:   m.PetrolTypeID,
		PetrolTypeName: m.PetrolType.Name, // struct field
		Notes:          m.Notes,
	}
}

func (u *fuelLogUsecase) Create(req dto.CreateFuelLogRequest) (*dto.FuelLogResponse, error) {
	log := &model.FuelLog{
		Date:          req.Date,
		PricePerLiter: req.PricePerLiter,
		TotalPaid:     req.TotalPaid,
		LitersFilled:  req.LitersFilled,
		KmStart:       req.KmStart,
		KmEnd:         req.KmEnd,
		LocationID:    req.LocationID,
		PetrolTypeID:  req.PetrolTypeID,
		Notes:         req.Notes,
	}

	if err := u.repo.Create(log); err != nil {
		return nil, err
	}

	return u.toResponse(log), nil
}

func (u *fuelLogUsecase) GetAll() ([]dto.FuelLogResponse, error) {
	logs, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.FuelLogResponse
	for _, l := range logs {
		result = append(result, *u.toResponse(&l))
	}
	return result, nil
}

func (u *fuelLogUsecase) GetByID(id uint) (*dto.FuelLogResponse, error) {
	log, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return u.toResponse(log), nil
}

func (u *fuelLogUsecase) Update(id uint, req dto.UpdateFuelLogRequest) (*dto.FuelLogResponse, error) {
	log, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	log.Date = req.Date
	log.PricePerLiter = req.PricePerLiter
	log.TotalPaid = req.TotalPaid
	log.LitersFilled = req.LitersFilled
	log.KmStart = req.KmStart
	log.KmEnd = req.KmEnd
	log.LocationID = req.LocationID
	log.PetrolTypeID = req.PetrolTypeID
	log.Notes = req.Notes

	if err := u.repo.Update(log); err != nil {
		return nil, err
	}

	return u.toResponse(log), nil
}

func (u *fuelLogUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
