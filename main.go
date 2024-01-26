package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	author string
	title string
	publisher string
}

type Repository struct {
	DB *gorm.DB
}

func (r Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	// these paths dont apply RESTful principles
	api.Post("/create-book", r.CreateBook)
	api.Delete("/delete-book/:id", r.DeleteBook)
	api.Get("/get-books/:id", r.GetBookById)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load database")
	}

	r := Repository{db}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")

}
