// THIS IS A EXAMPLE FILE

package repository

import (
	"github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/models"
	"github.com/jinzhu/gorm"
)

// UserRepository queries data source
type HelloRepository interface {
	FindAll() ([]*models.Hello, error)
}

type helloRepository struct {
	Repository
}

// NewUserRepository factory
func NewHelloRepository(db *gorm.DB) *helloRepository {
	return &helloRepository{
		Repository: Repository{
			db: db,
		},
	}
}

func (r *helloRepository) FindAll() ([]*models.Hello, error) {
	var hellos []*models.Hello
	err := r.getDB().Find(&hellos).Error
	if err != nil {
		return nil, err
	}
	return hellos, nil
}
