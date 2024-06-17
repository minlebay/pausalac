package customer

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
)

// CustomerController handles customer-related endpoints
type CustomerController struct {
	Service *usecases.EntityService[domain.Customer, domain.NewCustomer]
}

// GetAll godoc
// @Summary Get all customers
// @Description Get all customers
// @Tags customers
// @Produce json
// @Success 200 {array} CustomerResponse
// @Failure 500 {object} MessageResponse
// @Router /customers [get]
func (ctrl *CustomerController) GetAll(c *gin.Context) {
	customers, err := ctrl.Service.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponseArray(customers))
}

// GetByID godoc
// @Summary Get a customer by ID
// @Description Get a customer by ID
// @Tags customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} CustomerResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /customers/{id} [get]
func (ctrl *CustomerController) GetByID(c *gin.Context) {
	id := c.Param("id")
	customer, err := ctrl.Service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(customer))
}

// Create godoc
// @Summary Create a new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body CreateCustomerRequest true "Create Customer"
// @Success 201 {object} CustomerResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /customers [post]
func (ctrl *CustomerController) Create(c *gin.Context) {
	var req CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateCreateCustomerRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err := ctrl.Service.Create(context.Background(), ToDomain(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ToResponse(customer))
}

// Update godoc
// @Summary Update a customer
// @Description Update an existing customer by ID
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param customer body UpdateCustomerRequest true "Update Customer"
// @Success 200 {object} CustomerResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /customers/{id} [put]
func (ctrl *CustomerController) Update(c *gin.Context) {
	id := c.Param("id")
	var req UpdateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateUpdateCustomerRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customerMap := ToDomainUpdate(&req)
	customer, err := ctrl.Service.Update(context.Background(), id, customerMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(customer))
}

// Delete godoc
// @Summary Delete a customer
// @Description Delete a customer by ID
// @Tags customers
// @Param id path string true "Customer ID"
// @Success 204 {object} nil
// @Failure 500 {object} MessageResponse
// @Router /customers/{id} [delete]
func (ctrl *CustomerController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.Service.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
