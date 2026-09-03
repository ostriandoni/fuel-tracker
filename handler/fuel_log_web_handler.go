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
	usecase usecase.FuelLogUsecase
}

func NewFuelLogWebHandler(u usecase.FuelLogUsecase) *FuelLogWebHandler {
	return &FuelLogWebHandler{usecase: u}
}

func (h *FuelLogWebHandler) Index(c *gin.Context) {
	logs, _ := h.usecase.GetAll()
	c.HTML(http.StatusOK, "layout.html", logs)
}

func (h *FuelLogWebHandler) parseForm(c *gin.Context) dto.CreateFuelLogRequest {
	date, _ := time.Parse("2006-01-02", c.PostForm("date"))
	price, _ := decimal.NewFromString(c.PostForm("price_per_liter"))
	paid, _ := decimal.NewFromString(c.PostForm("total_paid"))
	liters, _ := decimal.NewFromString(c.PostForm("liters_filled"))
	kmStart, _ := strconv.Atoi(c.PostForm("km_start"))
	kmEnd, _ := strconv.Atoi(c.PostForm("km_end"))

	return dto.CreateFuelLogRequest{
		Date: date, PricePerLiter: price, TotalPaid: paid, LitersFilled: liters,
		KmStart: kmStart, KmEnd: kmEnd,
		Location: c.PostForm("location"), Notes: c.PostForm("notes"),
	}
}

func (h *FuelLogWebHandler) Create(c *gin.Context) {
	req := h.parseForm(c)
	h.usecase.Create(req)

	logs, _ := h.usecase.GetAll()
	c.HTML(http.StatusOK, "table", logs)
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
		Location: req.Location, Notes: req.Notes,
	}

	log, _ := h.usecase.Update(uint(id), updateReq)
	c.HTML(http.StatusOK, "row", log)
}

func (h *FuelLogWebHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.usecase.Delete(uint(id))
	c.String(http.StatusOK, "")
}
