package database

import (
	"fmt"
	"graph/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=jportal port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}

	err = db.Migrator().AutoMigrate(&models.Company{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}
	return db, nil

}

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error in hashing the password : %w", err)
	}
	return string(hashedPass), nil

}

func CheckHashedPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
