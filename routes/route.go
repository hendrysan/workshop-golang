package routes

import (
	"net/http"
	"workshop1/controllers"
	"workshop1/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	apiVersion := router.Group("api/v1")
	apiVersion.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	//  apiVersion.POST("/login", authController.Login)

	FinancialController := controllers.NewFinancialController(db)
	apiVersion.POST("/login", FinancialController.Login)

	financialGroup := apiVersion.Group("/financial")
	financialGroup.Use(middlewares.AuthJWT())
	{
		financialGroup.POST("/", FinancialController.CreateFinancial)
		financialGroup.GET("/", FinancialController.GetAllFinancial)
		financialGroup.GET("/:id", FinancialController.GetFinancialById)
		financialGroup.PUT("/:id", FinancialController.UpdateFinancial)
		financialGroup.DELETE("/:id", FinancialController.DeleteFinancial)
	}
}
