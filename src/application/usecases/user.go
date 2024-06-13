package usecases

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	"time"
)

// PaginationResultUser is the structure for pagination result of user
type PaginationResultUser struct {
	Data       []domain.User
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}

type UserService struct {
	Repo *repo.UserRepository
}

func (s *UserService) GetAll(ctx context.Context) (*[]domain.User, error) {
	users, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return s.arrayToDomainMapper(users), nil
}

func (s *UserService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(user), nil
}

func (s *UserService) Create(ctx context.Context, newUser *domain.NewUser) (*domain.User, error) {
	user := s.fromDomainMapper(newUser)
	createdUser, err := s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(createdUser), nil
}

func (s *UserService) CreateAdmin(ctx context.Context, newUser *domain.NewUser) (*domain.User, error) {
	users, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if users != nil && len(*users) > 0 {
		return nil, domain.NewAppErrorWithType(domain.ResourceAlreadyExists)
	}

	user := s.fromDomainMapper(newUser)
	createdUser, err := s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(createdUser), nil
}

func (s *UserService) GetOneByMap(ctx context.Context, userMap map[string]interface{}) (*domain.User, error) {
	user, err := s.Repo.GetOneByMap(ctx, userMap)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(user), nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id string, userMap map[string]interface{}) (*domain.User, error) {
	updatedUser, err := s.Repo.Update(ctx, id, userMap)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(updatedUser), nil
}

// Мапперы
func (s *UserService) fromDomainMapper(newUser *domain.NewUser) *domain.User {
	return &domain.User{
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

func (s *UserService) toDomainMapper(user *domain.User) *domain.User {
	return &domain.User{
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

func (s *UserService) arrayToDomainMapper(users *[]domain.User) *[]domain.User {
	var domainUsers []domain.User
	for _, user := range *users {
		domainUsers = append(domainUsers, *s.toDomainMapper(&user))
	}
	return &domainUsers
}
