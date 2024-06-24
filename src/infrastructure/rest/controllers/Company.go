package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
	"time"
)

type CompanyController struct {
	Service *usecases.EntityService[domain.Company]
}

// GetAll godoc
// @Summary Get all companies
// @Description Get all companies
// @Tags companies
// @Produce json
// @Success 200 {array} domain.Company
// @Failure 500 {object} controllers.MessageResponse
// @Router /companies [get]
func (ctrl *CompanyController) GetAll(ctx *gin.Context) {
	companies, err := ctrl.Service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if companies == nil {
		for _, company := range companies {
			if company.BankAccounts == nil {
				company.BankAccounts = []*domain.BankAccount{}
			}
		}
	}
	ctx.JSON(http.StatusOK, companies)
}

// GetById godoc
// @Summary Get a company by ID
// @Description Get a company by ID
// @Tags companies
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} domain.Company
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /companies/{id} [get]
func (ctrl *CompanyController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	company, err := ctrl.Service.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if company.BankAccounts == nil {
		company.BankAccounts = []*domain.BankAccount{}
	}
	ctx.JSON(http.StatusOK, company)
}

// Create godoc
// @Summary Create a new company
// @Description Create a new company
// @Tags companies
// @Accept json
// @Produce json
// @Param company body domain.Company true "Create Company"
// @Success 201 {object} domain.Company
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /companies [post]
func (ctrl *CompanyController) Create(ctx *gin.Context) {
	var c domain.Company

	c.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	c.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	author, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	c.Author = author.(string)

	if err := ctx.ShouldBindJSON(&c); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Id = primitive.NewObjectID()

	company, err := ctrl.Service.Create(ctx, &c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, company)
}

// Update godoc
// @Summary Update a company
// @Description Update a company
// @Tags companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Param company body domain.Company true "Update Company"
// @Success 200 {object} domain.Company
// @Failure 400 {object} controllers.MessageResponse
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /companies/{id} [put]
func (ctrl *CompanyController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var c domain.Company
	if err := ctx.ShouldBindJSON(&c); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	company, err := ctrl.Service.Update(ctx, id, &c)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, company)
}

// Delete godoc
// @Summary Delete a company
// @Description Delete a company
// @Tags companies
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} controllers.MessageResponse
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /companies/{id} [delete]
func (ctrl *CompanyController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ctrl.Service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, MessageResponse{Message: "Company deleted successfully"})
}
