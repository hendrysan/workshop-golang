package controllers

import (
	"net/http"
	"workshop1/helpers"
	"workshop1/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FinancialController struct {
	DB *gorm.DB
}

func NewFinancialController(db *gorm.DB) *FinancialController {
	return &FinancialController{DB: db}
}

func (controller *FinancialController) CreateFinancial(ctx *gin.Context) {

	var body models.FinancialCreateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.JSON(http.StatusBadRequest, helpers.Response(http.StatusBadRequest, "Bad Request", err.Error(), nil))
		return
	}

	financial := models.Financial{
		Category:    body.Category,
		Nominal:     body.Nominal,
		Description: body.Description,
	}

	err = controller.DB.Create(&financial).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response(http.StatusOK, "OK", "Success create data", financial))
}

func (controller *FinancialController) GetAllFinancial(ctx *gin.Context) {
	search := ctx.Query("search")
	category := ctx.Query("category")

	queryBase := controller.DB
	if search != "" {
		queryBase = queryBase.Where("description LIKE ?", "%"+search+"%")
	}

	if category != "" {
		queryBase = queryBase.Where("category", category)
	}

	var financials []models.Financial
	err := queryBase.Find(&financials).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response(http.StatusOK, "OK", "Success get data", financials))

}

func (controller *FinancialController) GetFinancialById(ctx *gin.Context) {

	id := ctx.Param("id")

	var financial models.Financial
	err := controller.DB.First(&financial, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response(http.StatusOK, "OK", "Success get data", financial))
}

func (controller *FinancialController) UpdateFinancial(ctx *gin.Context) {

	id := ctx.Param("id")

	var body models.FinancialCreateRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response(http.StatusBadRequest, "Bad Request", err.Error(), nil))
		return
	}

	var financial models.Financial
	err = controller.DB.First(&financial, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil))
		return
	}

	financial.Category = body.Category
	financial.Nominal = body.Nominal
	financial.Description = body.Description

	err = controller.DB.Save(&financial).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response(http.StatusOK, "OK", "Success update data", financial))
}

func (controller *FinancialController) DeleteFinancial(ctx *gin.Context) {

	id := ctx.Param("id")

	var financial models.Financial
	err := controller.DB.First(&financial, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil))
		return
	}

	err = controller.DB.Delete(&financial).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.Response(http.StatusOK, "OK", "Success delete data", financial))
}
