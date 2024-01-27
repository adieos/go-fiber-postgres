package main

// TODO: UPDATE LISTENER AND REFACTOR IF POSSIBLE !!!

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
	log.Println("Sukses nambah")
	return nil
}

func (r *Repository) ModifyBookHandler(ctx *fiber.Ctx) error {
	bookModel := model.Books{}
	book := Book{}
	id := ctx.Params("id")
	err := ctx.BodyParser(&book)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			fiber.Map{"message": "Unable to process response"})
		log.Println("Error 422: ", err)
		return err
	}

	result := r.db.Model(&bookModel).Where("id = ?", id).Updates(book)
	if result.Error != nil {
		ctx.Status(http.StatusInternalServerError).JSON(
			fiber.Map{"message": "Unable to update book"})
		log.Println("Error 500: ", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		ctx.Status(http.StatusNotFound).JSON(
			fiber.Map{"message": "Unable to update book: ID not found"})
		log.Println("Error 404: id gaada")
		return nil
	}

	ctx.Status(http.StatusOK).JSON(
		fiber.Map{"message": "book edited successfully"})
	log.Println("Sukses edit")
	return nil
}

func (r *Repository) DeleteBookHandler(ctx *fiber.Ctx) error {
	bookModel := model.Books{}
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Unable to delete book: ID cannot be empty"})
		log.Println("Error 400: id kosong")
		return nil
	}

	result := r.db.Delete(bookModel, id)
	if result.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Unable to delete book"})
		log.Println("Error 400: ", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		ctx.Status(http.StatusNotFound).JSON(
			fiber.Map{"message": "Unable to delete book: ID not found"})
		log.Println("Error 404: ID ga ada")
		return nil
	}

	ctx.Status(http.StatusOK).JSON(
		fiber.Map{"message": "Book deleted successfully"})
	log.Println("Sukses hapus")
	return nil

}

func (r *Repository) GetAllBooksHandler(ctx *fiber.Ctx) error {
	books := []model.Books{}

	err := r.db.Find(&books).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Unable to fetch books"})
		log.Println("Error 400: ", err)
		return err
	}

	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Books fetched successfully",
		"data":    books,
	})
	log.Println("Sukses get all")
	return nil

}

func (r *Repository) GetBookByIDHandler(ctx *fiber.Ctx) error {
	book := model.Books{}
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "ID cannot be empty"})
		log.Println("Error 400: ID kosong")
		return nil
	}

	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.Status(http.StatusNotFound).JSON(
				fiber.Map{"message": "Unable to fetch book: ID not found"})
			log.Println("Error 404: ID gaada")
			return err
		}
		ctx.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Unable to fetch book"})
		log.Println("Error 400: ", err)
		return err
	}

	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book fetched successfully",
		"data":    book,
	})
	log.Println("Sukses ambil buku")
	return nil
}

func (r *Repository) SetRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", r.CreateBookHandler)
	api.Put("/books/:id", r.ModifyBookHandler)
	api.Delete("/books/:id", r.DeleteBookHandler)
	api.Get("/books", r.GetAllBooksHandler)
	api.Get("/books/:id", r.GetBookByIDHandler)
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
