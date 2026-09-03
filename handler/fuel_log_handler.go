package handler

import (
	"net/http"
	"strconv"

	"fuel-tracker/dto"
	"fuel-tracker/helper"
	"fuel-tracker/usecase"

	"github.com/gin-gonic/gin"
)

type FuelLogHandler struct {
	usecase usecase.FuelLogUsecase
}

func NewFuelLogHandler(u usecase.FuelLogUsecase) *FuelLogHandler {
	return &FuelLogHandler{usecase: u}
}

func (h *FuelLogHandler) Create(c *gin.Context) {
	var req dto.CreateFuelLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.usecase.Create(req)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, "fuel log created", result)
}

func (h *FuelLogHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "success", result)
}

func (h *FuelLogHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	result, err := h.usecase.GetByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "fuel log not found")
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "success", result)
}

func (h *FuelLogHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.UpdateFuelLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.usecase.Update(uint(id), req)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "fuel log updated", result)
}

func (h *FuelLogHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusOK, "fuel log deleted", nil)
}
