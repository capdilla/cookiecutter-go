// THIS IS A EXAMPLE FILE

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Hello this is the description for this model
//
// swagger:model merchant
type Hello struct {
	BaseModel

	// The name for this record
	//
	// required: true
	Name string `json:"name" validate:"required"`
}

// TableName sets Merchant's table name to `merchant`
func (hello Hello) TableName() string {
	return "hello"
}

// BeforeCreate sets the creation date and a new ID
func (hello *Hello) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

// BeforeUpdate sets the date of the latest update
func (hello *Hello) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
