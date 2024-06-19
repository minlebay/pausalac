package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	usecases "pausalac/src/application/usecases"
)

// UserController is a controller for managing users
type UserController struct {
	Service *usecases.UserService
}

// GetAll godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} UserResponse
// @Failure 500 {object} MessageResponse
// @Router /users [get]
func (ctrl *UserController) GetAll(c *gin.Context) {
	users, err := ctrl.Service.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToDomainArray(users))
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get a single user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse
// @Failure 404 {object} MessageResponse
// @Router /users/{id} [get]
func (ctrl *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.Service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, ToResponse(user))
}

// CreateAdmin godoc
// @Summary Create a new admin
// @Description CreateAdmin method creates a new admin user if there are no users in the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "New User"
// @Success 201 {object} UserResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users/createadmin [post]
func (ctrl *UserController) CreateAdmin(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateCreateUserRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin user already exists"})
		return
	}
	req.Author = author.(string)

	userDomain := ToDomain(&req)
	user, err := ctrl.Service.CreateAdmin(context.Background(), userDomain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ToResponse(user))
}

// Create godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "New User"
// @Success 201 {object} UserResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [post]
func (ctrl *UserController) Create(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateCreateUserRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin user already exists"})
		return
	}
	req.Author = author.(string)

	userDomain := ToDomain(&req)
	user, err := ctrl.Service.Create(context.Background(), userDomain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ToResponse(user))
}

// Update godoc
// @Summary Update an existing user
// @Description Update an existing user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "Update User"
// @Success 200 {object} UserResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users/{id} [put]
func (ctrl *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ValidateUpdateUserRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userMap := ToDomainUpdate(&req)
	user, err := ctrl.Service.Update(context.Background(), id, userMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ToResponse(user))
}

// Delete godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Param id path string true "User ID"
// @Success 204 {object} nil
// @Failure 500 {object} MessageResponse
// @Router /users/{id} [delete]
func (ctrl *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.Service.Delete(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
