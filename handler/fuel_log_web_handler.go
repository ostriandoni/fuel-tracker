package handler

import (
	"net/http"
	"strconv"
	"time"

	"fuel-tracker/dto"
	"fuel-tracker/usecase"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type FuelLogWebHandler struct {
	usecase           usecase.FuelLogUsecase
	locationUsecase   usecase.LocationUsecase
	petrolTypeUsecase usecase.PetrolTypeUsecase
}

func NewFuelLogWebHandler(u usecase.FuelLogUsecase, lu usecase.LocationUsecase, ptu usecase.PetrolTypeUsecase) *FuelLogWebHandler {
	return &FuelLogWebHandler{usecase: u, locationUsecase: lu, petrolTypeUsecase: ptu}
}

func (h *FuelLogWebHandler) Index(c *gin.Context) {
	logs, _ := h.usecase.GetAll()
	c.HTML(http.StatusOK, "layout.html", logs)
}

func (h *FuelLogWebHandler) parseForm(c *gin.Context) dto.CreateFuelLogRequest {
	date, _ := time.Parse("2006-01-02", c.PostForm("date"))
	paid, _ := decimal.NewFromString(c.PostForm("total_paid"))
	liters, _ := decimal.NewFromString(c.PostForm("liters_filled"))
	kmStart, _ := strconv.Atoi(c.PostForm("km_start"))
	locationID, _ := strconv.Atoi(c.PostForm("location_id"))
	petrolTypeID, _ := strconv.Atoi(c.PostForm("petrol_type_id"))

	price := decimal.Zero
	if pt, err := h.petrolTypeUsecase.GetByID(uint(petrolTypeID)); err == nil {
		price = pt.Price
	}

	return dto.CreateFuelLogRequest{
		Date: date, PricePerLiter: price, TotalPaid: paid, LitersFilled: liters,
		KmStart: kmStart, LocationID: uint(locationID), PetrolTypeID: uint(petrolTypeID),
		Notes: c.PostForm("notes"),
	}
}

func (h *FuelLogWebHandler) Create(c *gin.Context) {
	req := h.parseForm(c)
	if _, err := h.usecase.Create(req); err != nil {
		c.String(http.StatusBadRequest, "Create failed: "+err.Error())
		return
	}

	logs, _ := h.usecase.GetAll()
	c.Writer.WriteHeader(http.StatusOK)
	c.HTML(http.StatusOK, "table", logs)
	c.Writer.Write([]byte(`<div id="modal-container" hx-swap-oob="true"></div>`))
}

func (h *FuelLogWebHandler) EditForm(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "not found")
		return
	}
	c.HTML(http.StatusOK, "edit_row", log)
}

func (h *FuelLogWebHandler) ViewRow(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log, _ := h.usecase.GetByID(uint(id))
	c.HTML(http.StatusOK, "row", log)
}

func (h *FuelLogWebHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	req := h.parseForm(c)

	updateReq := dto.UpdateFuelLogRequest{
		Date: req.Date, PricePerLiter: req.PricePerLiter, TotalPaid: req.TotalPaid,
		LitersFilled: req.LitersFilled, KmStart: req.KmStart, KmEnd: req.KmEnd,
		LocationID: req.LocationID, PetrolTypeID: req.PetrolTypeID, Notes: req.Notes,
	}

	log, err := h.usecase.Update(uint(id), updateReq)
	if err != nil {
		c.String(http.StatusBadRequest, "Update failed: "+err.Error())
		return
	}
	c.HTML(http.StatusOK, "row", log)
}

func (h *FuelLogWebHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.usecase.Delete(uint(id))
	c.String(http.StatusOK, "")
}

func (h *FuelLogWebHandler) NewForm(c *gin.Context) {
	locations, _ := h.locationUsecase.GetAll()
	c.HTML(http.StatusOK, "create_form", gin.H{
		"Locations": locations,
		"Today":     time.Now().Format("2006-01-02"),
	})
}

func (h *FuelLogWebHandler) CloseModal(c *gin.Context) {
	c.String(http.StatusOK, "")
}
