package service

import "time"

// ServiceResponse defines the response payload for a service
type ServiceResponse struct {
	Id        string    `json:"id"`
	Author    string    `json:"author" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MessageResponse defines a generic response message
type MessageResponse struct {
	Message string `json:"message"`
}
