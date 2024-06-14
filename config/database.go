package config

import (
	"backend-platform/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres dbname=backend-platform password=root123 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Migrator().DropTable(&models.Follow{})
	if err = db.AutoMigrate(&models.Follow{}, &models.Discussion{}, &models.Comment{}, &models.Reply{}, &models.User{}); err != nil {
		return db, err
	}
	return db, nil
}

func GetDB() *gorm.DB {
	return DB
}
