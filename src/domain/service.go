package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    string             `bson:"user_id" binding:"required"`
	Name      string             `bson:"name" binding:"required"`
	Unit      string             `bson:"unit" binding:"required"`
	Price     int64              `bson:"price" binding:"required"`
	Quantity  int64              `bson:"quantity" binding:"required"`
	Total     int64              `bson:"total" binding:"required"`
	Create_at primitive.DateTime `bson:"create_at"`
	Update_at primitive.DateTime `bson:"update_at"`
}

type ServiceArray []Service

type NewService struct {
	UserId   string
	Name     string
	Unit     string
	Price    int64
	Quantity int64
	Total    int64
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
		Unit:      newService.Unit,
		Price:     newService.Price,
		Quantity:  newService.Quantity,
		Total:     newService.Total,
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
