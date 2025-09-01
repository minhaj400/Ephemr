package database

import (
	"fmt"
	"log"
	"time"

	user "github.com/Minhajxdd/Ephemr/internal/user/model"
	"github.com/Minhajxdd/Ephemr/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Cfg.DB_HOST,
		config.Cfg.DB_USER,
		config.Cfg.DB_PWD,
		config.Cfg.DB_DATABASE,
		config.Cfg.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error Connecting To Database")
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
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

	DB = db
}
