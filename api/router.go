package api

import (
	"pos-is-backend/api/middleware"
	"pos-is-backend/internal/domain/repository"
	"pos-is-backend/internal/handler"
	"pos-is-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Cashier
	cashierRepo := repository.NewCashierRepository(db)
	cashierService := service.NewCashierService(cashierRepo)
	cashierHandler := handler.NewCashierHandler(cashierService)

	cashier := router.Group("cashiers")
	{
		cashier.POST("", cashierHandler.CreateCashier)
		cashier.GET("", cashierHandler.GetAllCashier)
		cashier.GET("/:cashierId", cashierHandler.GetDetailCashier)
		cashier.PUT("/:cashierId", cashierHandler.UpdateCashier)
		cashier.DELETE("/:cashierId", cashierHandler.DeleteCashier)
	}

	router.StaticFile("/", "static/index.html")

	return router
}
