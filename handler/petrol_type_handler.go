package handler

import (
	"net/http"
	"strconv"

	"fuel-tracker/helper"
	"fuel-tracker/usecase"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type PetrolTypeHandler struct {
	usecase usecase.PetrolTypeUsecase
}

func NewPetrolTypeHandler(u usecase.PetrolTypeUsecase) *PetrolTypeHandler {
	return &PetrolTypeHandler{usecase: u}
}

type createPetrolTypeRequest struct {
	LocationID uint            `json:"location_id" binding:"required"`
	Name       string          `json:"name" binding:"required"`
	Price      decimal.Decimal `json:"price" binding:"required"`
}

func (h *PetrolTypeHandler) Create(c *gin.Context) {
	var req createPetrolTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	pt, err := h.usecase.Create(req.LocationID, req.Name, req.Price)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusCreated, "petrol type created", pt)
}

func (h *PetrolTypeHandler) GetAll(c *gin.Context) {
	types, err := h.usecase.GetAll()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "success", types)
}

func (h *PetrolTypeHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	pt, err := h.usecase.GetByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "petrol type not found")
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "success", pt)
}

func (h *PetrolTypeHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req createPetrolTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	pt, err := h.usecase.Update(uint(id), req.LocationID, req.Name, req.Price)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "petrol type updated", pt)
}

func (h *PetrolTypeHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "petrol type deleted", nil)
}
