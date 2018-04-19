package entity

import (
	"errors"
	"strings"
)

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

func validateCategort(category string) error {
	var trimmedCategory = strings.Trim(category, " ")
	if trimmedCategory != "" && (trimmedCategory != "positive" && trimmedCategory != "negative") {
		return errors.New("category must have one of values: positive or negative")
	}
	return nil
}

func validateScore(score Score) error {
	if score < -1 || score > 100 {
		return errors.New("score is out of range:, it must be in 0-100")
	}
	return nil
}
