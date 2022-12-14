package models

import "time"

type User struct {
	ID          int64      `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber *string    `json:"phone_number"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Type        string     `json:"type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type CreateUserRequest struct {
	FirstName   string     `json:"first_name" binding:"required,min=2,max=30"`
	LastName    string     `json:"last_name" binding:"required,min=2,max=30"`
	PhoneNumber *string    `json:"phone_number"` // *
	Email       string     `json:"email" binding:"required,email"`
	Type        string     `json:"type" binding:"required,oneof=superadmin user partner"`
	Password    string     `json:"password" binding:"required,min=6,max=16"`
}

type UpdateUserRequest struct {
	FirstName   string     `json:"first_name" binding:"required,min=2,max=30"`
	LastName    string     `json:"last_name" binding:"required,min=2,max=30"`
	PhoneNumber *string    `json:"phone_number"` // *
	Email       string     `json:"email" binding:"required,email"`
	Type        string     `json:"type" binding:"required,oneof=superadmin user partner"`
	Password    string     `json:"password" binding:"required,min=6,max=16"`
}

type GetAllUsersResponse struct {
	Users []*User `json:"users"`
	Count int32   `json:"count"`
}

type VerifyRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}
