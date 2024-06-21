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

	// Convert invoices to SwaggerInvoice format
	var swaggerInvoices []domain.SwaggerInvoice
	for _, inv := range invoices {
		swaggerInvoices = append(swaggerInvoices, toSwaggerInvoice(*inv))
	}

	ctx.JSON(http.StatusOK, swaggerInvoices)
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

	ctx.JSON(http.StatusOK, toSwaggerInvoice(*invoice))
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
	var swaggerInvoice domain.SwaggerInvoice
	if err := ctx.ShouldBindJSON(&swaggerInvoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice := toDomainInvoice(swaggerInvoice)
	invoice.Id = primitive.NewObjectID()
	invoice.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	invoice.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	author, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	invoice.Author = author.(string)

	createdInvoice, err := ctrl.Service.Create(ctx, invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, toSwaggerInvoice(*createdInvoice))
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
	var swaggerInvoice domain.SwaggerInvoice
	if err := ctx.ShouldBindJSON(&swaggerInvoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice := toDomainInvoice(swaggerInvoice)
	invoice.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	updatedInvoice, err := ctrl.Service.Update(ctx, id, invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, toSwaggerInvoice(*updatedInvoice))
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

// Convert domain.Invoice to domain.SwaggerInvoice
func toSwaggerInvoice(invoice domain.Invoice) domain.SwaggerInvoice {
	return domain.SwaggerInvoice{
		Id:            invoice.Id.Hex(),
		Comment:       invoice.Comment,
		Number:        invoice.Number,
		TraidingPlace: invoice.TraidingPlace,
		Type:          invoice.Type,
		Author:        invoice.Author,
		Client:        invoice.Client,
		BankAccount:   invoice.BankAccount,
		Status:        invoice.Status,
		Services:      invoice.Services,
		PaidValue:     invoice.PaidValue,
		ValueInRSD:    invoice.ValueInRSD,
		Date:          invoice.Date.Time().Format(time.RFC3339),
		PaidDate:      invoice.PaidDate.Time().Format(time.RFC3339),
		SentDate:      invoice.SentDate.Time().Format(time.RFC3339),
		TradingDate:   invoice.TradingDate.Time().Format(time.RFC3339),
		CreatedAt:     invoice.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt:     invoice.UpdatedAt.Time().Format(time.RFC3339),
	}
}

// Convert domain.SwaggerInvoice to domain.Invoice
func toDomainInvoice(swaggerInvoice domain.SwaggerInvoice) *domain.Invoice {
	date, _ := time.Parse(time.RFC3339, swaggerInvoice.Date)
	paidDate, _ := time.Parse(time.RFC3339, swaggerInvoice.PaidDate)
	sentDate, _ := time.Parse(time.RFC3339, swaggerInvoice.SentDate)
	tradingDate, _ := time.Parse(time.RFC3339, swaggerInvoice.TradingDate)

	return &domain.Invoice{
		Comment:       swaggerInvoice.Comment,
		Number:        swaggerInvoice.Number,
		TraidingPlace: swaggerInvoice.TraidingPlace,
		Type:          swaggerInvoice.Type,
		Author:        swaggerInvoice.Author,
		Client:        swaggerInvoice.Client,
		BankAccount:   swaggerInvoice.BankAccount,
		Status:        swaggerInvoice.Status,
		Services:      swaggerInvoice.Services,
		PaidValue:     swaggerInvoice.PaidValue,
		ValueInRSD:    swaggerInvoice.ValueInRSD,
		Date:          primitive.NewDateTimeFromTime(date),
		PaidDate:      primitive.NewDateTimeFromTime(paidDate),
		SentDate:      primitive.NewDateTimeFromTime(sentDate),
		TradingDate:   primitive.NewDateTimeFromTime(tradingDate),
	}
}
