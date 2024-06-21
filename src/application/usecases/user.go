package usecases

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
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
	EntityService EntityService[domain.User]
	Repo          repo.UserRepository
}

func (s *UserService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return s.EntityService.GetAll(ctx)
}

func (s *UserService) GetById(ctx context.Context, id string) (*domain.User, error) {
	return s.EntityService.GetById(ctx, id)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.Repo.GetByEmail(ctx, email)
}

func (s *UserService) Create(ctx context.Context, newUser *domain.User) (*domain.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.HashPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	newUser.HashPassword = string(hashPassword)

	return s.EntityService.Create(ctx, newUser)
}

// CreateAdmin creates a new admin user if there are no users in the database
func (s *UserService) CreateAdmin(ctx context.Context, newUser *domain.User) (*domain.User, error) {
	users, err := s.EntityService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if users != nil && len(users) > 0 {
		return nil, domain.NewAppErrorWithType(domain.ResourceAlreadyExists)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.HashPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	newUser.HashPassword = string(hashPassword)

	return s.EntityService.Create(ctx, newUser)
}

func (s *UserService) Update(ctx context.Context, id string, user *domain.User) (*domain.User, error) {
	if user.HashPassword != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashPassword), bcrypt.DefaultCost)
		if err != nil {
			return nil, domain.NewAppErrorWithType(domain.UnknownError)
		}
		user.HashPassword = string(hashPassword)
	}

	return s.EntityService.Update(ctx, id, user)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	users, err := s.EntityService.GetAll(ctx)
	if err != nil {
		return err
	}

	if users != nil && len(users) == 1 {
		return domain.NewAppErrorWithType("Cannot delete the last user")
	}

	return s.EntityService.Delete(ctx, id)
}
