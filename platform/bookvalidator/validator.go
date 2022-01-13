package bookvalidator

import (
	"time"
	"w-r-api/platform"
)

// simple non-informative validator
func IsValid(book platform.Book) bool {
	if len(book.Id) < 25 || len(book.Id) > 70 {
		return false
	}
	if len(book.Name) < 2 || len(book.Name) > 70 {
		return false
	}
	if len(book.Author) < 2 || len(book.Author) > 70 {
		return false
	}
	if len(book.Isbn) < 10 || len(book.Isbn) > 20 {
		return false
	}
	if len(book.Publisher) < 2 || len(book.Publisher) > 70 {
		return false
	}
	if len(book.Genre) < 2 || len(book.Genre) > 50 {
		return false
	}
	if book.YearOfPublication < 0 || book.YearOfPublication > time.Now().Year()+1 {
		return false
	}
	if book.Pages < 1 || book.Pages > 20000 {
		return false
	}

	return true
}
