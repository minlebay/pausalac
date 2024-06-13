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

type CustomerRepository struct {
	Collection *mongo.Collection
}

func (r *CustomerRepository) GetAll(ctx context.Context) (*[]domain.Customer, error) {
	var customers []domain.Customer
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	if err := cursor.All(ctx, &customers); err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &customers, nil
}

func (r *CustomerRepository) Create(ctx context.Context, customer *domain.Customer) (*domain.Customer, error) {
	_, err := r.Collection.InsertOne(ctx, customer)
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
	return customer, nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, id string) (*domain.Customer, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	var customer domain.Customer
	err = r.Collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewAppErrorWithType(domain.NotFound)
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &customer, nil
}

func (r *CustomerRepository) Update(ctx context.Context, id string, customerMap map[string]interface{}) (*domain.Customer, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	customerMap["updated_at"] = primitive.NewDateTimeFromTime(time.Now())
	update := bson.M{"$set": customerMap}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var customer domain.Customer
	err = r.Collection.FindOneAndUpdate(ctx, bson.M{"_id": objId}, update, opts).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewAppErrorWithType(domain.NotFound)
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &customer, nil
}

func (r *CustomerRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.NewAppErrorWithType(domain.UnknownError)
	}
	res, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return domain.NewAppErrorWithType(domain.UnknownError)
	}
	if res.DeletedCount == 0 {
		return domain.NewAppErrorWithType(domain.NotFound)
	}
	return nil
}
