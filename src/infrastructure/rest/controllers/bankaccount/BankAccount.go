package bankaccount

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
)

// BankAccountController handles bank account-related endpoints
type BankAccountController struct {
	Service *usecases.EntityService[domain.BankAccount, domain.NewBankAccount]
}

// GetAll godoc
// @Summary Get all bank accounts
// @Description Get all bank accounts
// @Tags bankaccounts
// @Produce json
// @Success 200 {array} BankAccountResponse
// @Failure 500 {object} MessageResponse
// @Router /bankaccounts [get]
func (ctrl *BankAccountController) GetAll(c *gin.Context) {
	bankAccounts, err := ctrl.Service.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponseArray(bankAccounts))
}

// GetByID godoc
// @Summary Get a bank account by ID
// @Description Get a bank account by ID
// @Tags bankaccounts
// @Produce json
// @Param id path string true "Bank Account ID"
// @Success 200 {object} BankAccountResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /bankaccounts/{id} [get]
func (ctrl *BankAccountController) GetByID(c *gin.Context) {
	id := c.Param("id")
	bankAccount, err := ctrl.Service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(bankAccount))
}

// Create godoc
// @Summary Create a new bank account
// @Description Create a new bank account
// @Tags bankaccounts
// @Accept json
// @Produce json
// @Param bankaccount body CreateBankAccountRequest true "Create Bank Account"
// @Success 201 {object} BankAccountResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /bankaccounts [post]
func (ctrl *BankAccountController) Create(c *gin.Context) {
	var req CreateBankAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateCreateBankAccountRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bankAccount, err := ctrl.Service.Create(context.Background(), ToDomain(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ToResponse(bankAccount))
}

// Update godoc
// @Summary Update a bank account
// @Description Update a bank account
// @Tags bankaccounts
// @Accept json
// @Produce json
// @Param id path string true "Bank Account ID"
// @Param bankaccount body UpdateBankAccountRequest true "Update Bank Account"
// @Success 200 {object} BankAccountResponse
// @Failure 400 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /bankaccounts/{id} [put]
func (ctrl *BankAccountController) Update(c *gin.Context) {
	id := c.Param("id")
	var req UpdateBankAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateUpdateBankAccountRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bankAccount, err := ctrl.Service.Update(context.Background(), id, ToDomainUpdate(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(bankAccount))
}

// Delete godoc
// @Summary Delete a bank account
// @Description Delete a bank account
// @Tags bankaccounts
// @Produce json
// @Param id path string true "Bank Account ID"
// @Success 200 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /bankaccounts/{id} [delete]
func (ctrl *BankAccountController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.Service.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "Bank Account deleted successfully"})
}
