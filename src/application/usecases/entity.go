package usecases

import (
	"context"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
)

type PaginationResult[T domain.Types] struct {
	Data       []T
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}

type EntityService[T domain.Types] struct {
	Repo *repo.DefaultRepository[T]
}

func (s *EntityService[T]) GetAll(ctx context.Context) ([]*T, error) {
	entities, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *EntityService[T]) GetById(ctx context.Context, id string) (*T, error) {
	entity, err := s.Repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *EntityService[T]) Create(ctx context.Context, newEntity *T) (*T, error) {
	createdEntity, err := s.Repo.Create(ctx, newEntity)
	if err != nil {
		return nil, err
	}
	return createdEntity, nil
}

func (s *EntityService[T]) Update(ctx context.Context, id string, entity *T) (*T, error) {
	updatedEntity, err := s.Repo.Update(ctx, id, entity)
	if err != nil {
		return nil, err
	}
	return updatedEntity, nil
}

func (s *EntityService[T]) Delete(ctx context.Context, id string) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
