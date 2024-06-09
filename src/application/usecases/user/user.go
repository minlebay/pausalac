package user

import (
	"context"
	errs "github.com/minlebay/pausalac/src/domain/errors"
	domainUser "github.com/minlebay/pausalac/src/domain/user"
	userRepository "github.com/minlebay/pausalac/src/infrastructure/repository/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service struct {
	Repo *userRepository.Repository
}

func NewService(repo *userRepository.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetAll(ctx context.Context) (*[]domainUser.User, error) {
	users, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return s.arrayToDomainMapper(users), nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*domainUser.User, error) {
	user, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(user), nil
}

func (s *Service) Create(ctx context.Context, newUser *domainUser.NewUser) (*domainUser.User, error) {
	user := s.fromDomainMapper(newUser)
	createdUser, err := s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(createdUser), nil
}

func (s *Service) CreateAdmin(ctx context.Context, newUser *domainUser.NewUser) (*domainUser.User, error) {
	users, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if users != nil && len(*users) > 0 {
		return nil, errs.NewAppErrorWithType(errs.ResourceAlreadyExists)
	}

	user := s.fromDomainMapper(newUser)
	createdUser, err := s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(createdUser), nil
}

func (s *Service) GetOneByMap(ctx context.Context, userMap map[string]interface{}) (*domainUser.User, error) {
	user, err := s.Repo.GetOneByMap(ctx, userMap)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(user), nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, id string, userMap map[string]interface{}) (*domainUser.User, error) {
	updatedUser, err := s.Repo.Update(ctx, id, userMap)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(updatedUser), nil
}

// Мапперы
func (s *Service) fromDomainMapper(newUser *domainUser.NewUser) *domainUser.User {
	return &domainUser.User{
		ID:           primitive.NewObjectID(),
		UserName:     newUser.UserName,
		Email:        newUser.Email,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Status:       newUser.Status,
		Role:         newUser.Role,
		HashPassword: newUser.Password,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (s *Service) toDomainMapper(user *domainUser.User) *domainUser.User {
	return &domainUser.User{
		ID:           user.ID,
		UserName:     user.UserName,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Status:       user.Status,
		Role:         user.Role,
		HashPassword: user.HashPassword,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func (s *Service) arrayToDomainMapper(users *[]domainUser.User) *[]domainUser.User {
	var domainUsers []domainUser.User
	for _, user := range *users {
		domainUsers = append(domainUsers, *s.toDomainMapper(&user))
	}
	return &domainUsers
}
