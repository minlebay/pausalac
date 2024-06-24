package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceStatus string

const (
	PENDING   InvoiceStatus = "pending"
	PAID      InvoiceStatus = "paid"
	CANCELLED InvoiceStatus = "cancelled"
)

type Invoice struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Comment       string             `bson:"comment" json:"comment"`
	Number        string             `bson:"number" json:"number"`
	TraidingPlace string             `bson:"traiding_place" json:"traiding_place"`
	Type          string             `bson:"type" json:"type"`
	Author        string             `bson:"author" binding:"required" json:"author"`
	Client        Customer           `bson:"client" binding:"required" json:"client"`
	BankAccount   BankAccount        `bson:"bank_account" json:"bank_account"`
	Status        InvoiceStatus      `bson:"status" json:"status"`
	Services      []Service          `bson:"services" json:"services"`
	PaidValue     int64              `bson:"paid_value" json:"paid_value"`
	ValueInRSD    float64            `bson:"value_in_rsd" json:"value_in_rsd"`
	Date          primitive.DateTime `bson:"date" json:"date"`
	PaidDate      primitive.DateTime `bson:"paid_date" json:"paid_date"`
	SentDate      primitive.DateTime `bson:"sent_date" json:"sent_date"`
	TradingDate   primitive.DateTime `bson:"trading_date" json:"trading_date"`
	CreatedAt     primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt     primitive.DateTime `bson:"updated_at" json:"-"`
}

// SwaggerInvoice represents the invoice structure for Swagger documentation
// swagger:model
type SwaggerInvoice struct {
	Id            string        `json:"id" example:"5f8d04b2e8b2e7f8b2e8b2e8"`
	Comment       string        `json:"comment" example:"This is a comment"`
	Number        string        `json:"number" example:"INV-001"`
	TraidingPlace string        `json:"traiding_place" example:"Online"`
	Type          string        `json:"type" example:"Type1"`
	Author        string        `json:"author" example:"AuthorName"`
	Client        Customer      `json:"client"`
	BankAccount   BankAccount   `json:"bank_account"`
	Status        InvoiceStatus `json:"status"`
	Services      []Service     `json:"services"`
	PaidValue     int64         `json:"paid_value" example:"1000"`
	ValueInRSD    float64       `json:"value_in_rsd" example:"1000.50"`
	Date          string        `json:"date" swaggertype:"string" format:"date-time" example:"2022-03-07T13:45:00Z"`
	PaidDate      string        `json:"paid_date" swaggertype:"string" format:"date-time" example:"2022-03-07T13:45:00Z"`
	SentDate      string        `json:"sent_date" swaggertype:"string" format:"date-time" example:"2022-03-07T13:45:00Z"`
	TradingDate   string        `json:"trading_date" swaggertype:"string" format:"date-time" example:"2022-03-07T13:45:00Z"`
	CreatedAt     string        `json:"created_at" swaggertype:"string" format:"date-time" example:"2022-03-07T13:45:00Z"`
	UpdatedAt     string        `json:"updated_at" swaggertype:"string" format:"date-time" example:"2022-03-07T13:45:00Z"`
}

type InvoiceService interface {
	GetAll(context.Context) ([]*Invoice, error)
	GetById(ctx context.Context, id string) (*Invoice, error)
	Create(ctx context.Context, newInvoice *Invoice) (*Invoice, error)
	Update(ctx context.Context, id string, invoice *Invoice) (*Invoice, error)
	Delete(ctx context.Context, id string) error
}
