package user

import (
	"context"
	"errors"
	domainErrors "github.com/minlebay/pausalac/src/domain/errors"
	domainUser "github.com/minlebay/pausalac/src/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Repository is a struct that contains the database implementation for user entity
type Repository struct {
	Collection *mongo.Collection
}

// GetAll Fetch all user data
func (r *Repository) GetAll(ctx context.Context) (*[]domainUser.User, error) {
	var users []domainUser.User
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	return &users, nil
}

// Create ... Insert New data
func (r *Repository) Create(ctx context.Context, user *domainUser.User) (*domainUser.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.HashPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	user.HashPassword = string(pass)

	_, err = r.Collection.InsertOne(ctx, user)
	if err != nil {
		var mongoErr mongo.WriteException
		if errors.As(err, &mongoErr) {
			for _, writeErr := range mongoErr.WriteErrors {
				if writeErr.Code == 11000 {
					return nil, domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
				}
			}
		}
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	return user, nil
}

// GetOneByMap ... Fetch only one user by Map values
func (r *Repository) GetOneByMap(ctx context.Context, userMap map[string]interface{}) (*domainUser.User, error) {
	var user domainUser.User
	filter := bson.M{}
	for key, value := range userMap {
		filter[key] = value
	}
	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		}
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	return &user, nil
}

// GetByID ... Fetch only one user by ID
func (r *Repository) GetByID(ctx context.Context, id string) (*domainUser.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	var user domainUser.User
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		}
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	return &user, nil
}

// Update ... Update user
func (r *Repository) Update(ctx context.Context, id string, userMap map[string]interface{}) (*domainUser.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	userMap["updated_at"] = primitive.NewDateTimeFromTime(time.Now())
	pass, err := bcrypt.GenerateFromPassword([]byte(userMap["hash_password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	userMap["hash_password"] = string(pass)

	update := bson.M{"$set": userMap}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedUser domainUser.User
	err = r.Collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, opts).Decode(&updatedUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		}
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}

	return &updatedUser, nil
}

// Delete ... Delete user
func (r *Repository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	res, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}
	if res.DeletedCount == 0 {
		return domainErrors.NewAppErrorWithType(domainErrors.NotFound)
	}
	return nil
}
