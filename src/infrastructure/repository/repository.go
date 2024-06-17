package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pausalac/src/domain"
	"time"
)

type DefaultRepository[T domain.Types] struct {
	Collection *mongo.Collection
}

func (r *DefaultRepository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &entities, nil
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

func (r *DefaultRepository[T]) GetByID(ctx context.Context, id string) (*T, error) {
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

func (r *DefaultRepository[T]) Update(ctx context.Context, id string, entityMap map[string]interface{}) (*T, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	entityMap["updated_at"] = primitive.NewDateTimeFromTime(time.Now())
	update := bson.M{"$set": entityMap}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var entity T
	err = r.Collection.FindOneAndUpdate(ctx, bson.M{"_id": objId}, update, opts).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewAppErrorWithType(domain.NotFound)
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &entity, nil
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
