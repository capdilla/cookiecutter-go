package storage

import (
	"fmt"

	utils "github.com/DUNA-E-Commmerce/{{cookiecutter.app_name}}/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" /// just for dialect
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

// Init creates a connection to mysql database and
// migrates any new models
func Init(config *utils.Config) {
	fmt.Println("Connecting to database")
	dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err = gorm.Open("postgres", dbInfo)

	if err != nil {
		utils.Logger().Error("Failed to connect to databases: " + dbInfo)
		panic(err)
	}

	utils.Logger().Info("Database connected")
}

// GetDB gets db instance
func GetDB() *gorm.DB {
	return db
}

// CloseDB closes db connection
func CloseDB() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
