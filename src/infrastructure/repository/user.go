package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"pausalac/src/domain"
)

type UserRepository struct {
	DefaultRepository[domain.User]
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.NotFound)
	}
	return &user, nil

}
