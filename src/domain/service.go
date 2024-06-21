package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Author    string             `bson:"author" binding:"required" json:"author"`
	Name      string             `bson:"name" binding:"required" json:"name"`
	CreatedAt primitive.DateTime `bson:"create_at" json:"-"`
	UpdatedAt primitive.DateTime `bson:"update_at" json:"-"`
}

type SService interface {
	GetAll(ctx context.Context) ([]*Service, error)
	GetById(ctx context.Context, id string) (*Service, error)
	Create(ctx context.Context, newService *Service) (*Service, error)
	Update(ctx context.Context, id string, service *Service) (*Service, error)
	Delete(ctx context.Context, id string) error
}
