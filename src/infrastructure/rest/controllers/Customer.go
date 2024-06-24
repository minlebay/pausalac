package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
	"time"
)

type CustomerController struct {
	Service *usecases.EntityService[domain.Customer]
}

// GetAll godoc
// @Summary Get all customers
// @Description Get all customers
// @Tags customers
// @Produce json
// @Success 200 {array} domain.Customer
// @Failure 500 {object} controllers.MessageResponse
// @Router /customers [get]
func (ctrl *CustomerController) GetAll(ctx *gin.Context) {
	customers, err := ctrl.Service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

// GetById godoc
// @Summary Get a customer by ID
// @Description Get a customer by ID
// @Tags customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} domain.Customer
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /customers/{id} [get]
func (ctrl *CustomerController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	customer, err := ctrl.Service.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

// Create godoc
// @Summary Create a new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body domain.Customer true "Create Customer"
// @Success 201 {object} domain.Customer
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /customers [post]
func (ctrl *CustomerController) Create(ctx *gin.Context) {
	var customer domain.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.Id = primitive.NewObjectID()
	customer.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	customer.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	author, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	customer.Author = author.(string)

	c, err := ctrl.Service.Create(ctx, &customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, c)
}

// Update godoc
// @Summary Update a customer
// @Description Update an existing customer by ID
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param customer body domain.Customer true "Update Customer"
// @Success 200 {object} domain.Customer
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /customers/{id} [put]
func (ctrl *CustomerController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var customer domain.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	c, err := ctrl.Service.Update(ctx, id, &customer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, c)
}

// Delete godoc
// @Summary Delete a customer
// @Description Delete a customer by ID
// @Tags customers
// @Param id path string true "Customer ID"
// @Success 204 {object} nil
// @Failure 500 {object} controllers.MessageResponse
// @Router /customers/{id} [delete]
func (ctrl *CustomerController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ctrl.Service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, MessageResponse{Message: "Customer deleted successfully"})
}
