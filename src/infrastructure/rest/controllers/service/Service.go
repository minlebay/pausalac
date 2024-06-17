package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
)

// ServiceController handles service-related endpoints
type ServiceController struct {
	Service *usecases.EntityService[domain.Service, domain.NewService]
}

// GetAll godoc
// @Summary Get all services
// @Description Get all services
// @Tags services
// @Produce json
// @Success 200 {array} ServiceResponse
// @Failure 500 {object} MessageResponse
// @Router /services [get]
func (ctrl *ServiceController) GetAll(c *gin.Context) {
	services, err := ctrl.Service.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponseArray(services))
}

// GetByID godoc
// @Summary Get a service by ID
// @Description Get a service by ID
// @Tags services
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} ServiceResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /services/{id} [get]
func (ctrl *ServiceController) GetByID(c *gin.Context) {
	id := c.Param("id")
	service, err := ctrl.Service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(service))
}

// Create godoc
// @Summary Create a new service
// @Description Create a new service
// @Tags services
// @Accept json
// @Produce json
// @Param service body CreateServiceRequest true "Create Service"
// @Success 201 {object} ServiceResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /services [post]
func (ctrl *ServiceController) Create(c *gin.Context) {
	var req CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateCreateServiceRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service, err := ctrl.Service.Create(context.Background(), ToDomain(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ToResponse(service))
}

// Update godoc
// @Summary Update a service
// @Description Update a service
// @Tags services
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Param service body UpdateServiceRequest true "Update Service"
// @Success 200 {object} ServiceResponse
// @Failure 400 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /services/{id} [put]
func (ctrl *ServiceController) Update(c *gin.Context) {
	id := c.Param("id")
	var req UpdateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateUpdateServiceRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service, err := ctrl.Service.Update(context.Background(), id, ToDomainUpdate(&req))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(service))
}

// Delete godoc
// @Summary Delete a service
// @Description Delete a service
// @Tags services
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} MessageResponse
// @Failure 404 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /services/{id} [delete]
func (ctrl *ServiceController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.Service.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "Service deleted successfully"})
}
