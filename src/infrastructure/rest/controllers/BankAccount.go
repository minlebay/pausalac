package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
	"time"
)

type BankAccountController struct {
	Service *usecases.EntityService[domain.BankAccount]
}

// GetAll godoc
// @Summary Get all bank accounts
// @Description Get all bank accounts
// @Tags bankaccounts
// @Produce json
// @Success 200 {array} domain.BankAccount
// @Failure 500 {object} controllers.MessageResponse
// @Router /bankaccounts [get]
func (ctrl *BankAccountController) GetAll(ctx *gin.Context) {
	bankAccounts, err := ctrl.Service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bankAccounts)
}

// GetById godoc
// @Summary Get a bank account by ID
// @Description Get a bank account by ID
// @Tags bankaccounts
// @Produce json
// @Param id path string true "Bank Account ID"
// @Success 200 {object} domain.BankAccount
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /bankaccounts/{id} [get]
func (ctrl *BankAccountController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	bankAccount, err := ctrl.Service.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bankAccount)
}

// Create godoc
// @Summary Create a new bank account
// @Description Create a new bank account
// @Tags bankaccounts
// @Accept json
// @Produce json
// @Param bankaccount body domain.BankAccount true "Create Bank Account"
// @Success 201 {object} domain.BankAccount
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /bankaccounts [post]
func (ctrl *BankAccountController) Create(ctx *gin.Context) {
	var ba domain.BankAccount
	if err := ctx.ShouldBindJSON(&ba); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ba.Id = primitive.NewObjectID()
	ba.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	ba.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	author, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	ba.Author = author.(string)

	bankAccount, err := ctrl.Service.Create(ctx, &ba)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, bankAccount)
}

// Update godoc
// @Summary Update a bank account
// @Description Update a bank account
// @Tags bankaccounts
// @Accept json
// @Produce json
// @Param id path string true "Bank Account ID"
// @Param bankaccount body domain.BankAccount true "Update Bank Account"
// @Success 200 {object} domain.BankAccount
// @Failure 400 {object} controllers.MessageResponse
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /bankaccounts/{id} [put]
func (ctrl *BankAccountController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var ba domain.BankAccount
	if err := ctx.ShouldBindJSON(&ba); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ba.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	bankAccount, err := ctrl.Service.Update(ctx, id, &ba)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, bankAccount)
}

// Delete godoc
// @Summary Delete a bank account
// @Description Delete a bank account
// @Tags bankaccounts
// @Produce json
// @Param id path string true "Bank Account ID"
// @Success 200 {object} controllers.MessageResponse
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /bankaccounts/{id} [delete]
func (ctrl *BankAccountController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ctrl.Service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, MessageResponse{Message: "Bank account deleted successfully"})
}
