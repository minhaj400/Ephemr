package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *sql.DB

func ConnectDB() {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		Cfg.DB_HOST,
		Cfg.DB_USER,
		Cfg.DB_PWD,
		Cfg.DB_DATABASE,
		Cfg.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error Connecting To Database")
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error Connecting To Database")
	}

	fmt.Println("Connected to Database..")

	DB = sqlDB
}
