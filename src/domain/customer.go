package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerType string

const (
	INTERNAL CustomerType = "internal"
	FOREIGN  CustomerType = "foreign"
	INACTIVE CustomerType = "inactive"
)

type Customer struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Author             string             `bson:"author" binding:"required" json:"author"`
	Name               string             `bson:"name" binding:"required" json:"name"`
	TaxNumber          string             `bson:"tax_number" binding:"required" json:"tax_number"`
	RegistrationNumber string             `bson:"registration_number" binding:"required" json:"registration_number"`
	PhoneNumber        string             `bson:"phone_number" json:"phone_number"`
	Email              string             `bson:"email" json:"email"`
	Address            string             `bson:"address" json:"address"`
	City               string             `bson:"city" json:"city"`
	Country            string             `bson:"country" json:"country"`
	Currency           string             `bson:"currency" json:"currency"`
	CustomerType       CustomerType       `bson:"type" binding:"required" json:"customer_type"`
	CreatedAt          primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt          primitive.DateTime `bson:"updated_at" json:"-"`
}

type CustomerService interface {
	GetAll(context.Context) ([]*Customer, error)
	GetById(ctx context.Context, id string) (*Customer, error)
	Create(ctx context.Context, newCustomer *Customer) (*Customer, error)
	Update(ctx context.Context, id string, customer *Customer) (*Customer, error)
	Delete(ctx context.Context, id string) error
}
