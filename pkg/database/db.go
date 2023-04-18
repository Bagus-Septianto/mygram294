package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"mygram294/app/models"
	"mygram294/pkg/common"
)

var (
	db       *gorm.DB
	err      error
)

func StartDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", common.DBHost, common.DBUser,common.DBPassword, common.DBName, common.DBPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("connected to the database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.SocialMedia{}, models.Comment{})
	return db, err
}

func GetDB() *gorm.DB {
	return db
}