package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func Env_load() {
	if os.Getenv("ALPACA_ENV") == "" {
		os.Setenv("ALPACA_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("ALPACA_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Database() gorm.DB {
	Env_load()
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("USER_NAME"),
			os.Getenv("DATABASE_NAME"),
		))

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
