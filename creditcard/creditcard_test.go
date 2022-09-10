package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := "1234567890"
	cc, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}
	got := cc.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}

func TestInvalidReturnsError(t *testing.T) {
	want := "notanumber"
	_, err := creditcard.New(want)
	if err == nil {
		t.Fatal("expected error for invalid card number, got nil")
	}
}
