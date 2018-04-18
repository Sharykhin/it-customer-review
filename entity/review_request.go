package entity

import (
	"strings"

	"regexp"

	"github.com/pkg/errors"
)

// compile just once
var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

// ReviewRequest represents income body when a new review is going to be created
type ReviewRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

// Validate validates the current struct
func (rr ReviewRequest) Validate() error {

	var trimmedName, trimmedEmail, trimmedContent = strings.Trim(rr.Name, " "), strings.Trim(rr.Email, " "), strings.Trim(rr.Content, " ")
	if trimmedName == "" {
		return errors.New("name is required")
	}

	if len([]rune(trimmedName)) > 80 {
		return errors.New("name can not contain more than 80 characters")
	}

	if trimmedEmail == "" {
		return errors.New("email is required")
	}

	if len([]rune(trimmedEmail)) > 80 {
		return errors.New("email can not contain more than 80 characters")
	}

	if !re.MatchString(trimmedEmail) {
		return errors.New("enter a valid email address")
	}

	if trimmedContent == "" {
		return errors.New("content is required")
	}

	if len([]rune(trimmedContent)) > 2000 {
		return errors.New("content can not contain more than 2000 characters")
	}

	return nil
}
