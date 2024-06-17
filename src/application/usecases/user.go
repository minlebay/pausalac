package usecases

import (
	"context"
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
	Repo *repo.UserRepository
}

func (s *UserService) GetAll(ctx context.Context) (*[]domain.User, error) {
	users, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Create(ctx context.Context, newUser *domain.NewUser) (*domain.User, error) {
	user := newUser.ToDomainMapper()
	createdUser, err := s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserService) CreateAdmin(ctx context.Context, newUser *domain.NewUser) (*domain.User, error) {
	users, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if users != nil && len(*users) > 0 {
		return nil, domain.NewAppErrorWithType(domain.ResourceAlreadyExists)
	}

	user := newUser.ToDomainMapper()
	createdUser, err := s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserService) GetOneByMap(ctx context.Context, userMap map[string]interface{}) (*domain.User, error) {
	user, err := s.Repo.GetOneByMap(ctx, userMap)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id string, userMap map[string]interface{}) (*domain.User, error) {
	updatedUser, err := s.Repo.Update(ctx, id, userMap)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
