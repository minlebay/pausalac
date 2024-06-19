// Package auth provides the use case for authentication
package auth

import (
	"context"
	"errors"
	userRepository "pausalac/src/infrastructure/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
	"pausalac/src/application/security/jwt"
	errorsDomain "pausalac/src/domain"
)

// Auth contains the data of the authentication
type Auth struct {
	Email                     string
	AccessToken               string
	RefreshToken              string
	ExpirationAccessDateTime  time.Time
	ExpirationRefreshDateTime time.Time
}

// Service is a struct that contains the repository implementation for auth use case
type AuthService struct {
	UserRepository userRepository.UserRepository
}

// Login implements the login use case
func (s *AuthService) Login(ctx context.Context, user LoginUser) (*SecurityAuthenticatedUser, error) {
	userMap := map[string]any{"email": user.Email}
	domainUser, err := s.UserRepository.GetOneByMap(ctx, userMap)
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}
	if domainUser.Id.IsZero() {
		return &SecurityAuthenticatedUser{}, errorsDomain.NewAppError(errors.New("email or password does not match"), errorsDomain.NotAuthorized)
	}

	isAuthenticated := CheckPasswordHash(user.Password, domainUser.HashPassword)
	if !isAuthenticated {
		err = errorsDomain.NewAppError(err, errorsDomain.NotAuthorized)
		return &SecurityAuthenticatedUser{}, errorsDomain.NewAppError(errors.New("email or password does not match"), errorsDomain.NotAuthorized)
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(domainUser.Id, "access")
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}
	refreshTokenClaims, err := jwt.GenerateJWTToken(domainUser.Id, "refresh")
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}

	return secAuthUserMapper(domainUser, &Auth{
		AccessToken:               accessTokenClaims.Token,
		RefreshToken:              refreshTokenClaims.Token,
		ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
		ExpirationRefreshDateTime: refreshTokenClaims.ExpirationTime,
	}), err
}

// AccessTokenByRefreshToken implements the Access Token By Refresh Token use case
func (s *AuthService) AccessTokenByRefreshToken(ctx context.Context, refreshToken string) (*SecurityAuthenticatedUser, error) {
	claimsMap, err := jwt.GetClaimsAndVerifyToken(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	userMap := map[string]any{"id": claimsMap["id"]}
	domainUser, err := s.UserRepository.GetOneByMap(ctx, userMap)
	if err != nil {
		return nil, err

	}

	accessTokenClaims, err := jwt.GenerateJWTToken(domainUser.Id, "access")
	if err != nil {
		return &SecurityAuthenticatedUser{}, err
	}

	var expTime = int64(claimsMap["exp"].(float64))

	return secAuthUserMapper(domainUser, &Auth{
		AccessToken:               accessTokenClaims.Token,
		ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
		RefreshToken:              refreshToken,
		ExpirationRefreshDateTime: time.Unix(expTime, 0),
	}), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
