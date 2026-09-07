package handler

import (
	"net/http"
	"strconv"

	"fuel-tracker/helper"
	"fuel-tracker/model"
	"fuel-tracker/usecase"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	usecase usecase.LocationUsecase
}

func NewLocationHandler(u usecase.LocationUsecase) *LocationHandler {
	return &LocationHandler{usecase: u}
}

type createLocationRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *LocationHandler) Create(c *gin.Context) {
	var req createLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	loc, err := h.usecase.Create(req.Name)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(c, http.StatusCreated, "location created", loc)
}

func (h *LocationHandler) GetAll(c *gin.Context) {
	locations, err := h.usecase.GetAll()
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "success", locations)
}

func (h *LocationHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	loc, err := h.usecase.GetByID(uint(id))
	if err != nil {
		helper.ErrorResponse(c, http.StatusNotFound, "location not found")
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "success", loc)
}

func (h *LocationHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req createLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	loc, err := h.usecase.Update(uint(id), req.Name)
	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "location updated", loc)
}

func (h *LocationHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	helper.SuccessResponse(c, http.StatusOK, "location deleted", nil)
}

var _ = model.Location{} // keep import if unused elsewhere
