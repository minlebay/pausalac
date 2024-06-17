package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserName     string             `bson:"username"`
	Email        string             `bson:"email"`
	FirstName    string             `bson:"first_name"`
	LastName     string             `bson:"last_name"`
	Status       bool               `bson:"status"`
	Role         string             `bson:"role"`
	HashPassword string             `bson:"hash_password"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

type NewUser struct {
	UserName  string
	Email     string
	FirstName string
	LastName  string
	Role      string
	Password  string
	Status    bool
}

func (newUser *NewUser) ToDomainMapper() *User {
	return &User{
		ID:           primitive.NewObjectID(),
		UserName:     newUser.UserName,
		Email:        newUser.Email,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Status:       newUser.Status,
		Role:         newUser.Role,
		HashPassword: newUser.Password,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

type UserService interface {
	GetAll(context.Context) (*[]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, newUser *NewUser) (*User, error)
	GetOneByMap(ctx context.Context, userMap map[string]interface{}) (*User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, userMap map[string]interface{}) (*User, error)
}
