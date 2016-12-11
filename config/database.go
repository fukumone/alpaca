package config

import (
	"os"
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func env_load() {
	if os.Getenv("ALPACA_ENV") == "" {
		os.Setenv("ALPACA_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("ALPACA_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Database() *gorm.DB {
	env_load()

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER_NAME"),
		os.Getenv("DATABASE_USER_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	))

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
