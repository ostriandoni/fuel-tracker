package router

import (
	"fuel-tracker/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(fuelLogHandler *handler.FuelLogHandler) *gin.Engine {
	r := gin.Default()

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
