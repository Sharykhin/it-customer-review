package entity

import (
	"testing"
)

func TestNewReview(t *testing.T) {
	r := NewReview()
	if len(r.ID) != 36 {
		t.Fatalf("expected 38 characters but got %d", len(r.ID))
	}
}
