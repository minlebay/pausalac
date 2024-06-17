package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InvoiceStatus string

const (
	PENDING   InvoiceStatus = "pending"
	PAID      InvoiceStatus = "paid"
	CANCELLED InvoiceStatus = "cancelled"
)

type Invoice struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Cancelled     bool               `bson:"cancelled"`
	Comment       string             `bson:"comment"`
	Currency      string             `bson:"currency"`
	IBAN          string             `bson:"iban"`
	Number        string             `bson:"number"`
	SWIFT         string             `bson:"swift"`
	TraidingPlace string             `bson:"traiding_place"`
	Type          string             `bson:"type"`
	Author        User               `bson:"author" binding:"required"`
	Client        Customer           `bson:"client" binding:"required"`
	BankAccount   BankAccount        `bson:"bank_account"`
	Status        InvoiceStatus      `bson:"status"`
	Services      []Service          `bson:"services"`
	PaidValue     int64              `bson:"paid_value"`
	ValueInRSD    float64            `bson:"value_in_rsd"`
	CreatedAt     time.Time          `bson:"created_at"`
	Date          time.Time          `bson:"date"`
	PaidDate      time.Time          `bson:"paid_date"`
	SentDate      time.Time          `bson:"sent_date"`
	TradingDate   time.Time          `bson:"trading_date"`
	Created       time.Time          `bson:"created" binding:"required"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}

type NewInvoice struct {
	Author        User
	Created       time.Time
	BankAccount   BankAccount
	Cancelled     bool
	Client        Customer
	CreatedAt     time.Time
	Comment       string
	Currency      string
	IBAN          string
	Date          time.Time
	Number        string
	PaidDate      time.Time
	PaidValue     int64
	SentDate      time.Time
	Services      []Service
	Status        InvoiceStatus
	SWIFT         string
	TradingDate   time.Time
	TraidingPlace string
	Type          string
	ValueInRSD    float64
}

func (newInvoice *NewInvoice) ToDomainInvoiceMapper() *Invoice {
	return &Invoice{
		Id:            primitive.NewObjectID(),
		Author:        newInvoice.Author,
		Created:       newInvoice.Created,
		BankAccount:   newInvoice.BankAccount,
		Cancelled:     newInvoice.Cancelled,
		Client:        newInvoice.Client,
		Comment:       newInvoice.Comment,
		Currency:      newInvoice.Currency,
		IBAN:          newInvoice.IBAN,
		Date:          newInvoice.Date,
		Number:        newInvoice.Number,
		PaidDate:      newInvoice.PaidDate,
		PaidValue:     newInvoice.PaidValue,
		SentDate:      newInvoice.SentDate,
		Services:      newInvoice.Services,
		Status:        newInvoice.Status,
		SWIFT:         newInvoice.SWIFT,
		TradingDate:   newInvoice.TradingDate,
		TraidingPlace: newInvoice.TraidingPlace,
		Type:          newInvoice.Type,
		ValueInRSD:    newInvoice.ValueInRSD,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

type InvoiceService interface {
	GetAll(context.Context) (*[]Invoice, error)
	GetByID(ctx context.Context, id string) (*Invoice, error)
	Create(ctx context.Context, newInvoice *NewInvoice) (*Invoice, error)
	Update(ctx context.Context, id string, invoiceMap map[string]interface{}) (*Invoice, error)
	Delete(ctx context.Context, id string) error
}
