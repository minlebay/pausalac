package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Author    string             `bson:"author" binding:"required"`
	Name      string             `bson:"name" binding:"required"`
	Create_at primitive.DateTime `bson:"create_at"`
	Update_at primitive.DateTime `bson:"update_at"`
}

type ServiceArray []Service

type NewService struct {
	Author string
	Name   string
}

type NewServiceArray []NewService

func (newService NewServiceArray) ToDomainServiceArrayMapper() ServiceArray {
	var services []Service
	for _, service := range newService {
		s := service.ToDomainServiceMapper()
		services = append(services, *s)
	}
	return services
}

func (newService *NewService) ToDomainServiceMapper() *Service {
	return &Service{
		Id:        primitive.NewObjectID(),
		Name:      newService.Name,
		Create_at: primitive.NewDateTimeFromTime(time.Now()),
		Update_at: primitive.NewDateTimeFromTime(time.Now()),
	}
}

type SService interface {
	GetAll(ctx context.Context) (*[]Service, error)
	GetByID(ctx context.Context, id string) (*Service, error)
	Create(ctx context.Context, newService *NewService) (*Service, error)
	Update(ctx context.Context, id string, serviceMap map[string]interface{}) (*Service, error)
	Delete(ctx context.Context, id string) error
}
