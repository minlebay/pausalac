package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
	"time"
)

type ServiceController struct {
	Service *usecases.EntityService[domain.Service]
}

// GetAll godoc
// @Summary Get all services
// @Description Get all services
// @Tags services
// @Produce json
// @Success 200 {array} domain.Service
// @Failure 500 {object} controllers.MessageResponse
// @Router /services [get]
func (ctrl *ServiceController) GetAll(ctx *gin.Context) {
	services, err := ctrl.Service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, services)
}

// GetById godoc
// @Summary Get a service by ID
// @Description Get a service by ID
// @Tags services
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} domain.Service
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /services/{id} [get]
func (ctrl *ServiceController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	service, err := ctrl.Service.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, service)
}

// Create godoc
// @Summary Create a new service
// @Description Create a new service
// @Tags services
// @Accept json
// @Produce json
// @Param service body domain.Service true "Create Service"
// @Success 201 {object} domain.Service
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /services [post]
func (ctrl *ServiceController) Create(ctx *gin.Context) {
	var s domain.Service
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.Id = primitive.NewObjectID()
	s.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	s.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	author, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	s.Author = author.(string)

	service, err := ctrl.Service.Create(ctx, &s)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, service)
}

// Update godoc
// @Summary Update a service
// @Description Update a service
// @Tags services
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Param service body domain.Service true "Update Service"
// @Success 200 {object} domain.Service
// @Failure 400 {object} controllers.MessageResponse
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /services/{id} [put]
func (ctrl *ServiceController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var s domain.Service
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	service, err := ctrl.Service.Update(ctx, id, &s)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, service)
}

// Delete godoc
// @Summary Delete a service
// @Description Delete a service
// @Tags services
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} controllers.MessageResponse
// @Failure 404 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /services/{id} [delete]
func (ctrl *ServiceController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ctrl.Service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, MessageResponse{Message: "Service deleted successfully"})
}
