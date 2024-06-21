package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pausalac/src/domain"
)

type DefaultRepository[T domain.Types] struct {
	Collection *mongo.Collection
}

func (r *DefaultRepository[T]) GetAll(ctx context.Context) ([]*T, error) {
	var entities []*T
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return entities, nil
}

func (r *DefaultRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	_, err := r.Collection.InsertOne(ctx, entity)
	if err != nil {
		var mongoErr mongo.WriteException
		if errors.As(err, &mongoErr) {
			for _, writeErr := range mongoErr.WriteErrors {
				if writeErr.Code == 11000 {
					return nil, domain.NewAppErrorWithType(domain.ResourceAlreadyExists)
				}
			}
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return entity, nil
}

func (r *DefaultRepository[T]) GetById(ctx context.Context, id string) (*T, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	var entity T
	err = r.Collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewAppErrorWithType(domain.NotFound)
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &entity, nil
}

func (r *DefaultRepository[T]) Update(ctx context.Context, id string, entity *T) (*T, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}

	update := bson.M{"$set": entity}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := r.Collection.FindOneAndUpdate(ctx, bson.M{"_id": objId}, update, opts)
	if err = result.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.NewAppErrorWithType(domain.NotFound)
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}

	err = result.Decode(&entity)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return entity, nil
}

func (r *DefaultRepository[T]) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.NewAppErrorWithType(domain.UnknownError)
	}
	res, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return domain.NewAppErrorWithType(domain.UnknownError)
	}
	if res.DeletedCount == 0 {
		return domain.NewAppErrorWithType(domain.NotFound)
	}
	return nil
}
