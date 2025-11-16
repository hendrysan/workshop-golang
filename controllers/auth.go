package controllers

import (
	"net/http"
	"workshop1/helpers"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *FinancialController) Login(ctx *gin.Context) {
	var body LoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Response(http.StatusBadRequest, "Bad Request", err.Error(), nil))
		return
	}

	// TODO: ganti ini dengan cek user di DB
	if body.Username != "admin" || body.Password != "admin" {
		ctx.JSON(http.StatusUnauthorized, helpers.Response(http.StatusUnauthorized, "Unauthorized", "Invalid credentials", nil))
		return
	}

	token, err := helpers.GenerateToken(1) // contoh: userID = 1
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Response(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Response(http.StatusOK, "OK", "Login success", gin.H{"token": token}))
}
