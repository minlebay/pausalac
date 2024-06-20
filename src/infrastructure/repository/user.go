package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	domain "pausalac/src/domain"
	"time"
)

// UserRepository is a struct that contains the database implementation for user entity
type UserRepository struct {
	Collection *mongo.Collection
}

// GetAll Fetch all user data
func (r *UserRepository) GetAll(ctx context.Context) (*[]domain.User, error) {
	var users []domain.User
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &users, nil
}

// Create ... Insert New data
func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.HashPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	user.HashPassword = string(pass)

	_, err = r.Collection.InsertOne(ctx, user)
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
	return user, nil
}

// GetOneByMap ... Fetch only one user by Map values
func (r *UserRepository) GetOneByMap(ctx context.Context, userMap map[string]interface{}) (*domain.User, error) {
	var user domain.User
	filter := bson.M{}
	for key, value := range userMap {
		if key == "id" {
			objId, err := primitive.ObjectIDFromHex(value.(string))
			if err != nil {
				return nil, domain.NewAppErrorWithType(domain.UnknownError)
			}
			filter["_id"] = objId
			continue
		}

		filter[key] = value
	}
	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewAppErrorWithType(domain.NotFound)
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &user, nil
}

// GetByID ... Fetch only one user by ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	var user domain.User
	err = r.Collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewAppErrorWithType(domain.NotFound)
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	return &user, nil
}

// Update ... Update user
func (r *UserRepository) Update(ctx context.Context, id string, userMap map[string]interface{}) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	userMap["updated_at"] = primitive.NewDateTimeFromTime(time.Now())
	pass, err := bcrypt.GenerateFromPassword([]byte(userMap["hash_password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}
	userMap["hash_password"] = string(pass)

	update := bson.M{"$set": userMap}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedUser domain.User
	err = r.Collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, opts).Decode(&updatedUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.NewAppErrorWithType(domain.NotFound)
		}
		return nil, domain.NewAppErrorWithType(domain.UnknownError)
	}

	return &updatedUser, nil
}

// Delete ... Delete user
func (r *UserRepository) Delete(ctx context.Context, id string) error {
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
