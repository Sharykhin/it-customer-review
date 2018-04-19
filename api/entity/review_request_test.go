package entity

import (
	"sync"
	"testing"

	"github.com/pkg/errors"
)

func TestReviewRequest_Validate(t *testing.T) {

	tt := []struct {
		name        string
		request     ReviewRequest
		expectedErr error
	}{
		{
			name:        "name required",
			request:     ReviewRequest{Name: ""},
			expectedErr: errors.New("name is required"),
		},
		{
			name: "name is too long",
			request: ReviewRequest{Name: "aadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadas" +
				"aadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasds" +
				"adasaadfasdsadas"},
			expectedErr: errors.New("name can not contain more than 80 characters"),
		},
		{
			name:        "email required",
			request:     ReviewRequest{Name: "John", Email: ""},
			expectedErr: errors.New("email is required"),
		},
		{
			name: "email is too long",
			request: ReviewRequest{Name: "John", Email: "aadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadas" +
				"aadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasdsadasaadfasds" +
				"adasaadfasdsadas"},
			expectedErr: errors.New("email can not contain more than 80 characters"),
		},
		{
			name:        "email is invalid with no @",
			request:     ReviewRequest{Name: "John", Email: "a.com"},
			expectedErr: errors.New("enter a valid email address"),
		},
		{
			name:        "email is invalid with no domain",
			request:     ReviewRequest{Name: "John", Email: "a@com"},
			expectedErr: errors.New("enter a valid email address"),
		},
		{
			name:        "content is required",
			request:     ReviewRequest{Name: "John", Email: "john@john.com", Content: ""},
			expectedErr: errors.New("content is required"),
		},
	}

	var wg sync.WaitGroup
	for _, tc := range tt {
		wg.Add(1)
		go t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()
			err := tc.request.Validate()
			if err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("expected error %s, but got %s", tc.expectedErr.Error(), err.Error())
			} else if err != nil && tc.expectedErr == nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
	wg.Wait()
}
