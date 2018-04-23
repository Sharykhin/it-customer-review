package entity

import (
	"errors"
	"regexp"
	"strings"
)

// compile just once
var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func validateName(name string) error {
	var trimmedName = strings.Trim(name, " ")
	if trimmedName == "" {
		return errors.New("name is required")
	}

	if len([]rune(trimmedName)) > 80 {
		return errors.New("name can not contain more than 80 characters")
	}

	return nil
}

func validateEmail(email string) error {
	var trimmedEmail = strings.Trim(email, " ")
	if trimmedEmail == "" {
		return errors.New("email is required")
	}

	if len([]rune(trimmedEmail)) > 80 {
		return errors.New("email can not contain more than 80 characters")
	}

	if !re.MatchString(trimmedEmail) {
		return errors.New("enter a valid email address")
	}

	return nil
}

func validateContent(content string) error {
	var trimmedContent = strings.Trim(content, " ")
	if trimmedContent == "" {
		return errors.New("content is required")
	}

	if len([]rune(trimmedContent)) > 2000 {
		return errors.New("content can not contain more than 2000 characters")
	}

	return nil
}
