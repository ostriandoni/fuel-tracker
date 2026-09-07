package handler

import (
	"net/http"
	"strconv"

	"fuel-tracker/usecase"

	"github.com/gin-gonic/gin"
)

type LocationWebHandler struct {
	usecase usecase.PetrolTypeUsecase
}

func NewLocationWebHandler(u usecase.PetrolTypeUsecase) *LocationWebHandler {
	return &LocationWebHandler{usecase: u}
}

// GET /web/locations/:id/petrol-types
// Returns <option> list for the petrol type dropdown, filtered by location
func (h *LocationWebHandler) PetrolTypeOptions(c *gin.Context) {
	locationID, _ := strconv.Atoi(c.Param("id"))
	types, _ := h.usecase.GetByLocationID(uint(locationID))
	c.HTML(http.StatusOK, "petrol_type_options", types)
}
