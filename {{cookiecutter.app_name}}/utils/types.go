package utils

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Filter to pass on repository
type Filter struct {
	MerchantID *uuid.UUID
}

// EncryptMessageRequest message to be decrypted
type EncryptMessageRequest struct {
	Message string `json:"message"`
}

// LoginRequest type to login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// BaseModel gorm base model
type BaseModel struct {
	// The unique ID for this record
	//
	// required: true
	ID uint `json:"id" gorm:"primary_key"`

	// The creation time for this record
	//
	// required: true
	CreatedAt time.Time `json:"created_at"`

	// The modification time for this record
	//
	// required: true
	UpdatedAt time.Time `json:"last_updated"`

	// The deletion time for this record
	//
	// required: true
	DeletedAt *time.Time `json:"deleted_at"`
}

// UserResponse returned on fetching users
type UserResponse struct {
	BaseModel

	// The user email
	//
	// required: true
	Email string `json:"email"`

	// Role used to verify permissions
	//
	// required: true
	Role string `json:"role"`

	// The store code this driver belongs to
	//
	// required: true
	StoreCode string `json:"store_code"`

	// Name of the user
	//
	// required: true
	Name string `json:"name"`

	// If the user is a merchant-admin then this will be used to validate permissions
	//
	// required: false
	MerchantID *uuid.UUID `json:"merchant_id"`
}
