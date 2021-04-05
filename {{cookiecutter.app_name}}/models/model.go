package models

import "time"

// BaseModel reusable base model for the records.
//
// swagger:model baseModel
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

// BaseModelWithoutID reusable base model for the records.
// swagger:model baseModelWithoutID
type BaseModelWithoutID struct {
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
