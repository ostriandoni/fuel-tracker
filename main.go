package main

import (
	"log"
	"os"

	"fuel-tracker/config"
	"fuel-tracker/handler"
	"fuel-tracker/model"
	"fuel-tracker/repository"
	"fuel-tracker/router"
	"fuel-tracker/usecase"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system env")
	}

	db := config.ConnectDB()

	// auto migrate - remove in production, use proper migration tool instead
	db.AutoMigrate(&model.FuelLog{})

	fuelLogRepo := repository.NewFuelLogRepository(db)
	fuelLogUsecase := usecase.NewFuelLogUsecase(fuelLogRepo)
	fuelLogHandler := handler.NewFuelLogHandler(fuelLogUsecase)

	locationRepo := repository.NewLocationRepository(db)
	locationUsecase := usecase.NewLocationUsecase(locationRepo)
	locationHandler := handler.NewLocationHandler(locationUsecase)

	petrolTypeRepo := repository.NewPetrolTypeRepository(db)
	petrolTypeUsecase := usecase.NewPetrolTypeUsecase(petrolTypeRepo)
	petrolTypeHandler := handler.NewPetrolTypeHandler(petrolTypeUsecase)

	fuelLogWebHandler := handler.NewFuelLogWebHandler(fuelLogUsecase, locationUsecase, petrolTypeUsecase)
	locationWebHandler := handler.NewLocationWebHandler(petrolTypeUsecase)

	r := router.SetupRouter(fuelLogHandler, fuelLogWebHandler, locationWebHandler, locationHandler, petrolTypeHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
