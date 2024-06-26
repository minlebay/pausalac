package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	useCaseAuth "pausalac/src/application/usecases/auth"
	domainErrors "pausalac/src/domain"
	"pausalac/src/infrastructure/rest/controllers"
)

type Controller struct {
	AuthService useCaseAuth.AuthService
}

// Login godoc
// @Tags auth
// @Summary Login Email
// @Description Auth user by email and password
// @Param data body LoginRequest true "body data"
// @Success 200 {object} SecurityAuthenticatedUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /auth/login [post]
func (c *Controller) Login(ctx *gin.Context) {
	var request LoginRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	user := useCaseAuth.LoginUser{
		Email:    request.Email,
		Password: request.Password,
	}

	authDataUser, err := c.AuthService.Login(ctx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, authDataUser)
}

// GetAccessTokenByRefreshToken godoc
// @Tags auth
// @Summary GetAccessTokenByRefreshToken Email
// @Description Get access token by refresh token
// @Param data body AccessTokenRequest true "body data"
// @Success 200 {object} SecurityAuthenticatedUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /auth/access-token [post]
func (c *Controller) GetAccessTokenByRefreshToken(ctx *gin.Context) {
	var request AccessTokenRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	authDataUser, err := c.AuthService.AccessTokenByRefreshToken(ctx, request.RefreshToken)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, authDataUser)
}
