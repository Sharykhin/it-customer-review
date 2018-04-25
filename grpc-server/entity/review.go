package entity

import (
	"database/sql"

	"time"

	"database/sql/driver"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/grpc-server/util"
)

const (
	jsonTimeFormat = "2006-01-02T15:04:05"
)

type (
	// Review represents base entity for grpc server
	Review struct {
		ID        string
		Name      string
		Email     string
		Content   string
		Published bool
		Score     sql.NullInt64
		Category  sql.NullString
		CreatedAt JSONTime
		UpdatedAt JSONTime
	}
	// ReviewUpdate is wrapper around pb.ReviewUpdateRequest for easier validation
	ReviewUpdate struct {
		*pb.ReviewUpdateRequest
	}
	// JSONTime provides nullable time value in a preferable format
	JSONTime time.Time
)

//String returns datetime in a preferable format
func (t JSONTime) String() string {
	return time.Time(t).UTC().Format(jsonTimeFormat)
}

// Value returns string in UTC that should be saved in a database
func (t JSONTime) Value() (driver.Value, error) {
	return t.String(), nil
}

// NewReview is a fabric method that return a new review with generated uuid
func NewReview() *Review {

	uuid, err := util.NewUUID()
	if err != nil {
		// I intentionally run panic since if an error occurred it meant that something really hard went wrong
		panic("could not generate uuid")
	}
	return &Review{
		ID: uuid,
	}
}

// Validate fully validates review request
func (r Review) Validate() error {

	if err := validateName(r.Name); err != nil {
		return err
	}

	if err := validateEmail(r.Email); err != nil {
		return err
	}

	if err := validateContent(r.Content); err != nil {
		return err
	}

	if r.Category.Valid {
		if err := validateCategory(r.Category.String); err != nil {
			return err
		}
	}

	if r.Score.Valid {
		if err := validateScore(r.Score.Int64); err != nil {
			return err
		}
	}

	return nil
}

// Validate validates income request on update review
func (r ReviewUpdate) Validate() error {
	if !r.FieldsToUpdate.GetNameNull() {
		if err := validateName(r.FieldsToUpdate.GetNameValue()); err != nil {
			return err
		}
	}

	if !r.FieldsToUpdate.GetContentNull() {
		if err := validateContent(r.FieldsToUpdate.GetContentValue()); err != nil {
			return err
		}
	}

	if !r.FieldsToUpdate.GetScoreNull() {
		if err := validateScore(r.FieldsToUpdate.GetScoreValue()); err != nil {
			return err
		}
	}

	if !r.FieldsToUpdate.GetCategoryNull() {
		if err := validateCategory(r.FieldsToUpdate.GetCategoryValue()); err != nil {
			return err
		}
	}
	return nil
}
