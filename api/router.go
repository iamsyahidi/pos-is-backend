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
	verifyRepo := repository.NewVerifyRepository(db)
	verifyService := service.NewVerifyService(verifyRepo, cashierRepo)
	verifyHandler := handler.NewVerifyHandler(verifyService)

	// Category
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Product
	// Category
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	cashier := router.Group("cashiers")
	{
		cashier.POST("", cashierHandler.CreateCashier)
		cashier.GET("", cashierHandler.GetAllCashier)
		cashier.GET("/:cashierId", cashierHandler.GetDetailCashier)
		cashier.PUT("/:cashierId", cashierHandler.UpdateCashier)
		cashier.DELETE("/:cashierId", cashierHandler.DeleteCashier)
		cashier.GET("/:cashierId/passcode", verifyHandler.GetPasscode)
		cashier.POST("/:cashierId/login", verifyHandler.LoginPasscode)
		cashier.POST("/:cashierId/logout", verifyHandler.LogoutPasscode)
	}

	category := router.Group("categories")
	{
		category.POST("", categoryHandler.CreateCategory)
		category.GET("", categoryHandler.GetAllCategory)
		category.GET("/:categoryId", categoryHandler.GetDetailCategory)
		category.PUT("/:categoryId", categoryHandler.UpdateCategory)
		category.DELETE("/:categoryId", categoryHandler.DeleteCategory)
	}

	product := router.Group("products")
	{
		product.POST("", productHandler.CreateProduct)
		product.GET("", productHandler.GetAllProducts)
		product.GET("/:productId", productHandler.GetDetailProduct)
		product.PUT("/:productId", productHandler.UpdateProduct)
		product.DELETE("/:productId", productHandler.DeleteProduct)
	}

	router.StaticFile("/", "static/index.html")

	return router
}
