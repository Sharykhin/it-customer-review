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

func TestScore_Value(t *testing.T) {
	s := Score(-1)
	res, err := s.Value()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != nil {
		t.Fatalf("expected nil value but got: %v", res)
	}

	s = Score(10)
	res, err = s.Value()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != int64(10) {
		t.Fatalf("expected 10 value but got: %v", res)
	}
}
