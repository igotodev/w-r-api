package bookvalidator

import (
	"fmt"
	"time"
	"w-r-api/platform/entity"
)

// simple validator
func IsValid(book entity.Book) error {
	if len(book.Id) < 25 || len(book.Id) > 70 {
		return fmt.Errorf("id field length must be more than 25 characters and less than 70 characters ")
	}
	if len(book.Name) < 2 || len(book.Name) > 70 {
		return fmt.Errorf("name field length must be greater than 2 characters less than 70 characters")
	}
	if len(book.Author) < 2 || len(book.Author) > 70 {
		return fmt.Errorf("author field length must be greater than 2 characters less than 70 characters")
	}
	if len(book.Isbn) < 10 || len(book.Isbn) > 20 {
		return fmt.Errorf("isbn field length must be greater than 10 characters less than 20 characters")
	}
	if len(book.Publisher) < 2 || len(book.Publisher) > 70 {
		return fmt.Errorf("publisher field length must be greater than 2 characters less than 70 characters")
	}
	if len(book.Genre) < 2 || len(book.Genre) > 50 {
		return fmt.Errorf("genre field length must be greater than 2 characters less than 50 characters")
	}
	if book.YearOfPublication < 0 || book.YearOfPublication > time.Now().Year()+1 {
		return fmt.Errorf("year_of_publication value field must be greather than 0 and less than this year + 1")
	}
	if book.Pages < 1 || book.Pages > 10000 {
		return fmt.Errorf("pages value field must be greater than 0 and less than 10000 ")
	}

	return nil
}
