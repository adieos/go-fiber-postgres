package main

import (
	"log"
	"os"

	"github.com/adieos/go-fiber-postgres/model"
	"github.com/adieos/go-fiber-postgres/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	author    string
	title     string
	publisher string
}

type Repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// config
	config := &setup.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// connect and migrate db
	db, err := setup.NewConnection(config)
	if err != nil {
		log.Fatal("could not load database")
	}
	err := model.MigrateDB(db)
	if err != nil {
		log.Fatal("could not migrate database")
	}

	r := Repository{db}

	app := fiber.New()
	app.Listen(":8080")

}
