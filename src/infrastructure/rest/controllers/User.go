package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	usecases "pausalac/src/application/usecases"
	"pausalac/src/domain"
	"time"
)

type UserController struct {
	Service *usecases.UserService
}

// GetAll godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} domain.User
// @Failure 500 {object} controllers.MessageResponse
// @Router /users [get]
func (ctrl *UserController) GetAll(ctx *gin.Context) {
	users, err := ctrl.Service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// GetById godoc
// @Summary Get user by ID
// @Description Get a single user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} domain.User
// @Failure 404 {object} controllers.MessageResponse
// @Router /users/{id} [get]
func (ctrl *UserController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := ctrl.Service.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// CreateAdmin godoc
// @Summary Create a new admin
// @Description CreateAdmin method creates a new admin user if there are no users in the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "New User"
// @Success 201 {object} domain.User
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /users/createadmin [post]
func (ctrl *UserController) CreateAdmin(ctx *gin.Context) {
	var u domain.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Id = primitive.NewObjectID()
	u.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	u.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	author, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	u.Author = author.(string)

	user, err := ctrl.Service.CreateAdmin(ctx, &u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

// Create godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "New User"
// @Success 201 {object} domain.User
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /users [post]
func (ctrl *UserController) Create(ctx *gin.Context) {
	var u domain.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Id = primitive.NewObjectID()
	u.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	u.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	author, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Author not found"})
		return
	}
	u.Author = author.(string)

	user, err := ctrl.Service.Create(ctx, &u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

// Update godoc
// @Summary Update an existing user
// @Description Update an existing user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body domain.User true "Update User"
// @Success 200 {object} domain.User
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /users/{id} [put]
func (ctrl *UserController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var u domain.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	user, err := ctrl.Service.Update(ctx, id, &u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Param id path string true "User ID"
// @Success 204 {object} nil
// @Failure 500 {object} controllers.MessageResponse
// @Router /users/{id} [delete]
func (ctrl *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := ctrl.Service.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, MessageResponse{Message: "User deleted successfully"})
}
