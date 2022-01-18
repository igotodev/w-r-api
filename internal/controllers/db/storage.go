package db

import (
	"database/sql"
	"log"
	"w-r-api/internal/domain/book/model"
	"w-r-api/internal/domain/book/storage"
)

type bookStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) storage.Storage {
	return &bookStorage{db: db}
}

func (b bookStorage) GetByID(idStr string) *model.Book {
	rows, err := b.db.Query(`SELECT * FROM books WHERE id=$1`, idStr)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var book model.Book

	for rows.Next() {

		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Isbn, &book.Publisher, &book.Genre, &book.YearOfPublication, &book.Pages)
	}

	return &book
}

func (b bookStorage) GetAll() *[]model.Book {
	rows, err := b.db.Query(`SELECT * FROM books`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []model.Book

	for rows.Next() {
		var book model.Book
		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Isbn, &book.Publisher, &book.Genre, &book.YearOfPublication, &book.Pages)

		books = append(books, book)
	}

	return &books
}

func (b bookStorage) Create(book *model.Book) error {
	_, err := b.db.Exec(`INSERT INTO books (id, name, author, isbn, publisher, genre, year_of_publication, pages) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		book.Id, book.Name, book.Author, book.Isbn, book.Publisher, book.Genre, book.YearOfPublication, book.Pages)

	return err
}

func (b bookStorage) Update(book *model.Book, idStr string) error {
	_, err := b.db.Exec(`UPDATE books SET id=$2, name=$3, author=$4, isbn=$5, publisher=$6, genre=$7, year_of_publication=$8, pages=$9 WHERE id=$1`,
		idStr, book.Id, book.Name, book.Author, book.Isbn, book.Publisher, book.Genre, book.YearOfPublication, book.Pages)

	return err
}

func (b bookStorage) Delete(idStr string) error {
	_, err := b.db.Exec(`DELETE FROM books WHERE id=$1`, idStr)

	return err
}
