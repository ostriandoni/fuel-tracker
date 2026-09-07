package router

import (
	"fuel-tracker/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(fuelLogHandler *handler.FuelLogHandler, webHandler *handler.FuelLogWebHandler, locationWebHandler *handler.LocationWebHandler, locationHandler *handler.LocationHandler, petrolTypeHandler *handler.PetrolTypeHandler) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLFiles(
		"templates/layout.html",
		"templates/partials/table.html",
		"templates/partials/row.html",
		"templates/partials/edit_row.html",
		"templates/partials/create_form.html",
		"templates/partials/petrol_type_options.html",
	)
	r.Static("/static", "./static")

	r.GET("/", webHandler.Index)
	web := r.Group("/web/fuel-logs")
	{
		web.POST("", webHandler.Create)
		web.GET("/:id", webHandler.ViewRow)
		web.GET("/:id/edit", webHandler.EditForm)
		web.PUT("/:id", webHandler.Update)
		web.DELETE("/:id", webHandler.Delete)
		web.GET("/new", webHandler.NewForm)
		web.GET("/close-modal", webHandler.CloseModal)
	}

	r.GET("/web/locations/:id/petrol-types", locationWebHandler.PetrolTypeOptions)

	api := r.Group("/api/v1")
	{
		fuelLogs := api.Group("/fuel-logs")
		{
			fuelLogs.POST("", fuelLogHandler.Create)
			fuelLogs.GET("", fuelLogHandler.GetAll)
			fuelLogs.GET("/:id", fuelLogHandler.GetByID)
			fuelLogs.PUT("/:id", fuelLogHandler.Update)
			fuelLogs.DELETE("/:id", fuelLogHandler.Delete)
		}

		locations := api.Group("/locations")
		{
			locations.POST("", locationHandler.Create)
			locations.GET("", locationHandler.GetAll)
			locations.GET("/:id", locationHandler.GetByID)
			locations.PUT("/:id", locationHandler.Update)
			locations.DELETE("/:id", locationHandler.Delete)
		}

		petrolTypes := api.Group("/petrol-types")
		{
			petrolTypes.POST("", petrolTypeHandler.Create)
			petrolTypes.GET("", petrolTypeHandler.GetAll)
			petrolTypes.GET("/:id", petrolTypeHandler.GetByID)
			petrolTypes.PUT("/:id", petrolTypeHandler.Update)
			petrolTypes.DELETE("/:id", petrolTypeHandler.Delete)
		}
	}

	return r
}
