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

type EntityService[T domain.Types, N domain.NewTypes] struct {
	Repo       *repo.DefaultRepository[T]
	MapperFunc func(ctx context.Context, newEntity *N) *T
}

func (s *EntityService[T, N]) GetAll(ctx context.Context) (*[]T, error) {
	entities, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *EntityService[T, N]) GetByID(ctx context.Context, id string) (*T, error) {
	entity, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *EntityService[T, N]) Create(ctx context.Context, newEntity *N) (*T, error) {
	entity := s.MapperFunc(ctx, newEntity)
	createdEntity, err := s.Repo.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return createdEntity, nil
}

func (s *EntityService[T, N]) Update(ctx context.Context, id string, entityMap map[string]interface{}) (*T, error) {
	updatedEntity, err := s.Repo.Update(ctx, id, entityMap)
	if err != nil {
		return nil, err
	}
	return updatedEntity, nil
}

func (s *EntityService[T, N]) Delete(ctx context.Context, id string) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
