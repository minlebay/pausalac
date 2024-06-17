package service

import "time"

// ServiceResponse defines the response payload for a service
type ServiceResponse struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Unit      string    `json:"unit" binding:"required"`
	Price     int64     `json:"price" binding:"required"`
	Quantity  int64     `json:"quantity" binding:"required"`
	Total     int64     `json:"total" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MessageResponse defines a generic response message
type MessageResponse struct {
	Message string `json:"message"`
}
