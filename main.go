package main

// TODO: refactor it? idk :(

import (
	"log"
	"net/http"
	"os"

	"github.com/adieos/go-fiber-postgres/model"
	"github.com/adieos/go-fiber-postgres/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateBookHandler(ctx *fiber.Ctx) error {
	book := Book{}

	err := ctx.BodyParser(&book)
	if err != nil || book.Author == "" || book.Publisher == "" || book.Title == "" {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			fiber.Map{"message": "Unable process response"})
		log.Println("Error 422: ", err)
		return err
	}

	err = r.db.Create(&book).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Unable to create book"})
		log.Println("Error 400: ", err)
		return err
	}

	ctx.Status(http.StatusOK).JSON(
		fiber.Map{"message": "Book created successfully"})
	return nil
}

// func (r *Repository) PutBookListener(ctx *fiber.Ctx) error{

// }

func (r *Repository) SetRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", r.CreateBookHandler)
	// api.Put("/books/:id", r.PutBookListener)
	// api.Delete("/books/:id", r.DeleteBookListener)
	// api.Get("/books", r.GetAllBooksListener)
	// api.Get("/books/:id", r.GetBookByIDListener)
}

func main() {
	// load env variables
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
	err = model.MigrateDB(db)
	if err != nil {
		log.Fatal("could not migrate database")
	}

	r := Repository{db}
	app := fiber.New()
	r.SetRoutes(app) // NOTE: !!!
	app.Listen(":8080")

}
