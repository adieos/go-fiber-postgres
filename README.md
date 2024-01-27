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
2. Configure your PostgreSQL database settings in the `.env` file.
3. Run the application using `go run main.go`.

## API Endpoints

- `POST /api/books`: Add a new book to the database.
- `GET /api/books/:id`: Retrieve details of a specific book by ID.
- `GET /api/books`: Retrieve a list of all books.
- `PUT /api/books/:id`: Update details of a specific book by ID.
- `DELETE /api/books/:id`: Delete a specific book by ID.

## Example Usage

```bash
# Add a new book
curl -X POST -H "Content-Type: application/json" -d '{"title": "AC Forsaken", "author": "Anton Gill", "publisher": "Penguin Publishing"}' http://localhost:8080/api/books

# Retrieve book details by ID
curl -X GET http://localhost:8080/api/books/1
