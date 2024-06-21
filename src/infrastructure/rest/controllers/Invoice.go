package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
	"time"
)

// InvoiceController handles invoice-related endpoints
type InvoiceController struct {
	Service *usecases.EntityService[domain.Invoice]
}

// GetAll godoc
// @Summary Get all invoices
// @Description Get all invoices
// @Tags invoices
// @Produce json
// @Success 200 {array} domain.SwaggerInvoice
// @Failure 500 {object} controllers.MessageResponse
// @Router /invoices [get]
func (ctrl *InvoiceController) GetAll(ctx *gin.Context) {
	invoices, err := ctrl.Service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, invoices)
}

// GetById godoc
// @Summary Get an invoice by ID
// @Description Get an invoice by ID
// @Tags invoices
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} domain.SwaggerInvoice
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /invoices/{id} [get]
func (ctrl *InvoiceController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	invoice, err := ctrl.Service.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, invoice)
}

// Create godoc
// @Summary Create a new invoice
// @Description Create a new invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param invoice body domain.SwaggerInvoice true "Create Invoice"
// @Success 201 {object} domain.SwaggerInvoice
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /invoices [post]
func (ctrl *InvoiceController) Create(ctx *gin.Context) {
	var i domain.Invoice
	if err := ctx.ShouldBindJSON(&i); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	i.Id = primitive.NewObjectID()
	i.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	i.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	author, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	i.Author = author.(string)

	invoice, err := ctrl.Service.Create(ctx, &i)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, invoice)
}

// Update godoc
// @Summary Update an invoice
// @Description Update an invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Param invoice body domain.SwaggerInvoice true "Update Invoice"
// @Success 200 {object} domain.SwaggerInvoice
// @Failure 400 {object} controllers.MessageResponse
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /invoices/{id} [put]
func (ctrl *InvoiceController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var i domain.Invoice
	if err := ctx.ShouldBindJSON(&i); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	i.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	invoice, err := ctrl.Service.Update(ctx, id, &i)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, invoice)
}

// Delete godoc
// @Summary Delete an invoice
// @Description Delete an invoice
// @Tags invoices
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} controllers.MessageResponse
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /invoices/{id} [delete]
func (ctrl *InvoiceController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ctrl.Service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, MessageResponse{Message: "Invoice deleted successfully"})
}
