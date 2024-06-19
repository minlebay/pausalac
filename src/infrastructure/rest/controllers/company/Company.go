package company

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
)

// CompanyController handles company-related endpoints
type CompanyController struct {
	Service *usecases.EntityService[domain.Company, domain.NewCompany]
}

// GetAll godoc
// @Summary Get all companies
// @Description Get all companies
// @Tags companies
// @Produce json
// @Success 200 {array} CompanyResponse
// @Failure 500 {object} MessageResponse
// @Router /companies [get]
func (ctrl *CompanyController) GetAll(c *gin.Context) {
	companies, err := ctrl.Service.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if companies == nil {
		for _, company := range *companies {
			if company.BankAccounts == nil {
				company.BankAccounts = []domain.BankAccount{}
			}
		}
	}
	c.JSON(http.StatusOK, ToResponseArray(companies))
}

// GetByID godoc
// @Summary Get a company by ID
// @Description Get a company by ID
// @Tags companies
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} CompanyResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /companies/{id} [get]
func (ctrl *CompanyController) GetByID(c *gin.Context) {
	id := c.Param("id")
	company, err := ctrl.Service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if company.BankAccounts == nil {
		company.BankAccounts = []domain.BankAccount{}
	}
	c.JSON(http.StatusOK, ToResponse(company))
}

// Create godoc
// @Summary Create a new company
// @Description Create a new company
// @Tags companies
// @Accept json
// @Produce json
// @Param company body CreateCompanyRequest true "Create Company"
// @Success 201 {object} CompanyResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /companies [post]
func (ctrl *CompanyController) Create(c *gin.Context) {
	var req CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateCreateCompanyRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	req.Author = author.(string)

	company, err := ctrl.Service.Create(context.Background(), ToDomain(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ToResponse(company))
}

// Update godoc
// @Summary Update a company
// @Description Update a company
// @Tags companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Param company body UpdateCompanyRequest true "Update Company"
// @Success 200 {object} CompanyResponse
// @Failure 400 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /companies/{id} [put]
func (ctrl *CompanyController) Update(c *gin.Context) {
	id := c.Param("id")
	var req UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateUpdateCompanyRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company, err := ctrl.Service.Update(context.Background(), id, ToDomainUpdate(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(company))
}

// Delete godoc
// @Summary Delete a company
// @Description Delete a company
// @Tags companies
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /companies/{id} [delete]
func (ctrl *CompanyController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.Service.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "Company deleted successfully"})
}
