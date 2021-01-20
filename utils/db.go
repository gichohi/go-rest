package utils

import (
	"fmt"
	"github.com/gichohi/go-rest.git/models"
	"github.com/jinzhu/gorm"
	_ 	"github.com/jinzhu/gorm/dialects/postgres"
)

const (
	DB_HOST 	= "127.0.0.1"
	DB_PORT 	= "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = ""
	DB_NAME     = "users"
)


func GetDB() *gorm.DB {

	var db *gorm.DB
	var err error


	db, err = gorm.Open(
		"postgres",
		"host="+DB_HOST+" user="+DB_USER+
			" dbname="+DB_NAME+" sslmode=disable password="+DB_PASSWORD)

	if err != nil {
		fmt.Println("error", err)
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.User{})

	fmt.Println("Successfully connected!", db)
	return db
}
