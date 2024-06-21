package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"  json:"id"`
	Author       string             `bson:"author" binding:"required" json:"author"`
	UserName     string             `bson:"username" json:"username"`
	Email        string             `bson:"email" json:"email"`
	FirstName    string             `bson:"first_name" json:"first_name"`
	LastName     string             `bson:"last_name" json:"last_name"`
	Status       bool               `bson:"status" json:"status"`
	Role         string             `bson:"role" json:"role"`
	HashPassword string             `bson:"hash_password" json:"password"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt    primitive.DateTime `bson:"updated_at" json:"-"`
}

type UserService interface {
	GetAll(context.Context) ([]*User, error)
	GetById(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, newUser *User) (*User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, user *User) (*User, error)

	CreateAdmin(ctx context.Context, newUser *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
