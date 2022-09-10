package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmpopts"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "A Tale of Two Cities",
		Author: "Charles Dickens",
		Copies: 7,
	}
	want := 6
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "The Way of Kings",
		Author: "Brandon Sanderson",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	//if err != nil {
	//		t.Fatal(err)
	//}
	if !cmp.Equal(want, got,
		cmpopts.IgnoreUnexporeted(bookstore.Book())) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	want := bookstore.Book{
		ID:    2,
		Title: "The Power of Go: Tools",
	}
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: want,
	}
	got, err := catalog.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got,
		cmpopts.IgnoreUnexporeted(bookstore.Book())) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
	}
	_, err := catalog.GetBook(2)
	if err == nil {
		t.Error("error expected for nonexistent book ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:           "The Way of Kings",
		Author:          "Brandon Sanderson",
		Copies:          5,
		ID:              1,
		PriceCents:      3500,
		DiscountPercent: 15,
	}
	want := 2975
	got := book.NetPriceCents()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:      "The Way of Kings",
		Author:     "Brandon Sanderson",
		PriceCents: 3500,
	}
	want := 4000
	err := book.SetPriceCents(4000)
	if err != nil {
		t.Fatal(err)
	}
	got := book.PriceCents
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	err := book.SetPriceCents(-1)
	if err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{Title: "For the Love of Go"}
	want := "programming"
	err := book.SetCategory("programming")
	got := book.Category()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("expected category %q, got %q", want, got)
	}
}

func TestSetCategoryInvalidInput(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{Title: "For the Love of Go"}
	err := book.SetCategory("bogus")
	if err == nil {
		t.Fatal("want error fo rinvalid category, got nil")
	}
}
