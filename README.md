# Bookstore API

Welcome to the Bookstore API, a simple yet powerful solution for storing and retrieving book information. Built with Go and Fiber for the backend, and PostgreSQL as the database.

## Features

- **Store Books:** Add new books to the database with details like title, author, and publisher.
- **Retrieve Books:** Fetch information about stored books based on various criteria.
- **Powerful Backend:** Utilizes the Fiber framework for a fast and efficient API.

## Technologies Used

- Go
- Fiber (web framework for Go)
- PostgreSQL

## Getting Started

1. Clone the repository.
2. Configure your PostgreSQL database settings in the `config.yaml` file.
3. Run the application using `go run main.go`.

## API Endpoints

- `POST /api/books`: Add a new book to the database.
- `GET /api/books/:id`: Retrieve details of a specific book by ID.
- `GET /api/books`: Retrieve a list of all books.
- ... (add more endpoints as needed)

## Example Usage

```bash
# Add a new book
curl -X POST -H "Content-Type: application/json" -d '{"title": "Sample Book", "author": "John Doe", "publisher": "Sample Publisher"}' http://localhost:3000/api/books

# Retrieve book details by ID
curl http://localhost:3000/api/books/1
