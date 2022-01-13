package db

import (
	"database/sql"
	"fmt"
	"log"
	"w-r-api/config"
	"w-r-api/platform"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func OpenDB() {
	cfg := config.InitConfig()
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", cfg.Driver, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func SelectAll() []platform.Book {
	rows, err := db.Query(`SELECT * FROM books`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []platform.Book

	for rows.Next() {
		var book platform.Book
		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Isbn, &book.Publisher, &book.Genre, &book.YearOfPublication, &book.Pages)

		books = append(books, book)
	}

	return books
}

func Select(idStr string) platform.Book {
	rows, err := db.Query(`SELECT * FROM books WHERE id=$1`, idStr)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var book platform.Book

	for rows.Next() {

		rows.Scan(&book.Id, &book.Name, &book.Author, &book.Isbn, &book.Publisher, &book.Genre, &book.YearOfPublication, &book.Pages)
	}

	return book
}

func Insert(book platform.Book) {

	_, err := db.Exec(`INSERT INTO books (id, name, author, isbn, publisher, genre, year_of_publication, pages) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		book.Id, book.Name, book.Author, book.Isbn, book.Publisher, book.Genre, book.YearOfPublication, book.Pages)

	if err != nil {
		log.Println(err)
	}

}

func Update(book platform.Book, idStr string) {

	_, err := db.Exec(`UPDATE books SET id=$2, name=$3, author=$4, isbn=$5, publisher=$6, genre=$7, year_of_publication=$8, pages=$9 WHERE id=$1`,
		idStr, book.Id, book.Name, book.Author, book.Isbn, book.Publisher, book.Genre, book.YearOfPublication, book.Pages)

	if err != nil {
		log.Println(err)
	}

}

func Delete(idStr string) bool {

	_, err := db.Exec(`DELETE FROM books WHERE id=$1`, idStr)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
