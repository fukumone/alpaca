package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Env_load() {
	if os.Getenv("GIN_ENV") == "" {
		os.Setenv("GIN_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GIN_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Database() gorm.DB {
	Env_load()
	db, err := gorm.Open("postgres",
		fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable",
			os.Getenv("DATABASE_NAME"),
			os.Getenv("USER_NAME"),
			os.Getenv("PASSWORD"),
		))

	if err != nil {
		panic(err)
	}
	return db
}
