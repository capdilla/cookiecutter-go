package repository

import (
	storage "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/storage"
	"github.com/jinzhu/gorm"
)

// Repository the base repo
type Repository struct {
	db *gorm.DB
}

func (r *Repository) SetDB(db *gorm.DB) {
	r.db = db
}

func (r *Repository) getDB() *gorm.DB {
	if r.db != nil {
		return r.db
	}
	r.db = storage.GetDB()
	return r.db
}
