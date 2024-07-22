package db

import (
	"log"

	"github.com/shauryagupta3/recruitment-management-sys/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Job{})
	conn.AutoMigrate(&models.Profile{})
	return conn
}
