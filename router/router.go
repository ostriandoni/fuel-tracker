package router

import (
	"fuel-tracker/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(fuelLogHandler *handler.FuelLogHandler, webHandler *handler.FuelLogWebHandler) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLFiles(
		"templates/layout.html",
		"templates/partials/table.html",
		"templates/partials/row.html",
		"templates/partials/edit_row.html",
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
	}

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
	}

	return r
}
