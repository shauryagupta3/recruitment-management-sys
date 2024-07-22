package db

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shauryagupta3/recruitment-management-sys/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *pgxpool.Pool

func Init(url string) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	conn.AutoMigrate(&models.User{})
	return conn
}
