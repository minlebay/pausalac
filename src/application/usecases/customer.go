package usecases

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	"time"
)

type PaginationResultCustomer struct {
	Data       []domain.Customer
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}

type CustomerService struct {
	Repo *repo.CustomerRepository
}

func (s *CustomerService) GetAll(ctx context.Context) (*[]domain.Customer, error) {
	customers, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return s.arrayToDomainMapper(customers), nil
}

func (s *CustomerService) GetByID(ctx context.Context, id string) (*domain.Customer, error) {
	customer, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(customer), nil
}

func (s *CustomerService) Create(ctx context.Context, newCustomer *domain.NewCustomer) (*domain.Customer, error) {
	customer := s.fromDomainMapper(newCustomer)
	createdCustomer, err := s.Repo.Create(ctx, customer)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(createdCustomer), nil
}

func (s *CustomerService) Update(ctx context.Context, id string, customerMap map[string]interface{}) (*domain.Customer, error) {
	updatedCustomer, err := s.Repo.Update(ctx, id, customerMap)
	if err != nil {
		return nil, err
	}
	return s.toDomainMapper(updatedCustomer), nil
}

func (s *CustomerService) Delete(ctx context.Context, id string) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerService) fromDomainMapper(newCustomer *domain.NewCustomer) *domain.Customer {
	return &domain.Customer{
		ID:                 primitive.NewObjectID(),
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

func (s *CustomerService) toDomainMapper(customer *domain.Customer) *domain.Customer {
	return &domain.Customer{
		ID:                 customer.ID,
		Name:               customer.Name,
		TaxNumber:          customer.TaxNumber,
		RegistrationNumber: customer.RegistrationNumber,
		PhoneNumber:        customer.PhoneNumber,
		Email:              customer.Email,
		Address:            customer.Address,
		City:               customer.City,
		Country:            customer.Country,
		Currency:           customer.Currency,
		CustomerType:       customer.CustomerType,
		CreatedAt:          customer.CreatedAt,
		UpdatedAt:          customer.UpdatedAt,
	}
}

func (s *CustomerService) arrayToDomainMapper(customers *[]domain.Customer) *[]domain.Customer {
	var result []domain.Customer
	for _, customer := range *customers {
		result = append(result, *s.toDomainMapper(&customer))
	}
	return &result
}
