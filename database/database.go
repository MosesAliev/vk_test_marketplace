// database/database.go
package database

import (
	"fmt"
	"log"
	"os"
	"vk_test_marketplace/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	// создаём URL для соединения с базой данных.
	// Имя пользователя базы данных, пароль и имя базы данных
	// берутся из переменных окружения,
	// они описаны в файле .env
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	// ,
	// os.Getenv("DB_USER"),
	// os.Getenv("DB_PASSWORD"),
	// os.Getenv("DB_NAME"),

	// создаём подключение к базе данных.
	// В &gorm.Config настраивается логер,
	// который будет сохранять информацию
	// обо всех активностях с базой данных.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(1)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")
	db.AutoMigrate(model.User{}, model.Ad{})

	DB = Dbinstance{
		Db: db,
	}
}
