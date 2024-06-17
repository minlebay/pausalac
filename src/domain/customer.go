package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CustomerType string

const (
	INTERNAL CustomerType = "internal"
	FOREIGN  CustomerType = "foreign"
	INACTIVE CustomerType = "inactive"
)

type Customer struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name" binding:"required"`
	TaxNumber          string             `bson:"tax_number" binding:"required"`
	RegistrationNumber string             `bson:"registration_number" binding:"required"`
	PhoneNumber        string             `bson:"phone_number"`
	Email              string             `bson:"email"`
	Address            string             `bson:"address"`
	City               string             `bson:"city"`
	Country            string             `bson:"country"`
	Currency           string             `bson:"currency"`
	CustomerType       CustomerType       `bson:"type" binding:"required"` // internal, foreign, inactive
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
}

type NewCustomer struct {
	Name               string
	TaxNumber          string
	RegistrationNumber string
	PhoneNumber        string
	Email              string
	Address            string
	City               string
	Country            string
	Currency           string
	CustomerType       CustomerType
}

func (newCustomer *NewCustomer) ToDomainCustomerMapper() *Customer {
	return &Customer{
		Id:                 primitive.NewObjectID(),
		Name:               newCustomer.Name,
		TaxNumber:          newCustomer.TaxNumber,
		RegistrationNumber: newCustomer.RegistrationNumber,
		PhoneNumber:        newCustomer.PhoneNumber,
		Email:              newCustomer.Email,
		Address:            newCustomer.Address,
		City:               newCustomer.City,
		Country:            newCustomer.Country,
		Currency:           newCustomer.Currency,
		CustomerType:       newCustomer.CustomerType,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}

type CustomerService interface {
	GetAll(context.Context) (*[]Customer, error)
	GetByID(ctx context.Context, id string) (*Customer, error)
	Create(ctx context.Context, newCustomer *NewCustomer) (*Customer, error)
	Update(ctx context.Context, id string, customerMap map[string]interface{}) (*Customer, error)
	Delete(ctx context.Context, id string) error
}
