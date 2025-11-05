package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	var conn *gorm.DB

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Close(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		log.Println("error closing db:", err)
		return
	}
	sqlDB.Close()
}
