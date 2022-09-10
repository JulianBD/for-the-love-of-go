package bookstore

import (
	"errors"
	"fmt"
)

// Book struct represents a book object
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        string
}

// Catalog type represents a catalog of books implemented
// as map
type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies -= 1
	return b, nil
}

func (catalog Catalog) GetAllBooks() []Book {
	var result []Book
	for _, book := range catalog {
		result = append(result, book)
	}
	return result
}

func (catalog Catalog) GetBook(id int) (Book, error) {
	b, present := catalog[id]
	if !present {
		return b, fmt.Errorf("ID %d doesn't exist", id)
	}
	return b, nil
}

func (b Book) NetPriceCents() int {
	return (100 - b.DiscountPercent) * (b.PriceCents / 100)
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("negative price %d", price)
	}
	b.PriceCents = price
	return nil
}

func (b *Book) Category() string {
	return b.category
}

func (b *Book) SetCategory(category string) error {
	if category != "programming" {
		return t.Errorf("%q not a valid category", category)
	}
	b.category = category
	return nil
}
