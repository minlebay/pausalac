package invoice

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
)

// InvoiceController handles invoice-related endpoints
type InvoiceController struct {
	Service *usecases.EntityService[domain.Invoice, domain.NewInvoice]
}

// GetAll godoc
// @Summary Get all invoices
// @Description Get all invoices
// @Tags invoices
// @Produce json
// @Success 200 {array} InvoiceResponse
// @Failure 500 {object} MessageResponse
// @Router /invoices [get]
func (ctrl *InvoiceController) GetAll(c *gin.Context) {
	invoices, err := ctrl.Service.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponseArray(invoices))
}

// GetByID godoc
// @Summary Get an invoice by ID
// @Description Get an invoice by ID
// @Tags invoices
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} InvoiceResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /invoices/{id} [get]
func (ctrl *InvoiceController) GetByID(c *gin.Context) {
	id := c.Param("id")
	invoice, err := ctrl.Service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(invoice))
}

// Create godoc
// @Summary Create a new invoice
// @Description Create a new invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param invoice body CreateInvoiceRequest true "Create Invoice"
// @Success 201 {object} InvoiceResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /invoices [post]
func (ctrl *InvoiceController) Create(c *gin.Context) {
	var req CreateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateCreateInvoiceRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	req.Author = author.(string)

	invoice, err := ctrl.Service.Create(context.Background(), ToDomain(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ToResponse(invoice))
}

// Update godoc
// @Summary Update an invoice
// @Description Update an invoice
// @Tags invoices
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Param invoice body UpdateInvoiceRequest true "Update Invoice"
// @Success 200 {object} InvoiceResponse
// @Failure 400 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /invoices/{id} [put]
func (ctrl *InvoiceController) Update(c *gin.Context) {
	id := c.Param("id")
	var req UpdateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateUpdateInvoiceRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	invoice, err := ctrl.Service.Update(context.Background(), id, ToDomainUpdate(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(invoice))
}

// Delete godoc
// @Summary Delete an invoice
// @Description Delete an invoice
// @Tags invoices
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /invoices/{id} [delete]
func (ctrl *InvoiceController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.Service.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "Invoice deleted successfully"})
}
