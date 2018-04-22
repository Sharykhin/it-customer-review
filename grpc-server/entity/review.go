package entity

import (
	"regexp"

	"database/sql"

	"encoding/json"

	"database/sql/driver"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
	"github.com/Sharykhin/it-customer-review/grpc-server/util"
)

// compile just once
var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type (
	// Review represents base entity for grpc server
	Review struct {
		*pb.ReviewRequest
		ID        string
		CreatedAt JSONTime
		Category  sql.NullString
		Score     Score
	}

	ReviewM struct {
		ID        string         `json:"id"`
		Name      string         `json:"name"`
		Email     string         `json:"email"`
		Content   string         `json:"content"`
		Published sql.NullBool   `json:"published"`
		Score     sql.NullInt64  `json:"score"`
		Category  sql.NullString `json:"category"`
		CreatedAt string         `json:"created_at"`
	}

	// ReviewUpdate represent a private entity that would be used for updating an existing review
	ReviewUpdate struct {
		*pb.ReviewUpdateRequest
	}
	// Score is a specific type that converts into null in case -1 was provided
	Score int64
)

// MarshalJSON implements Marshaler interface to return null
func (s Score) MarshalJSON() ([]byte, error) {
	if s == -1 {
		return json.Marshal(nil)
	}

	return json.Marshal(int64(s))
}

// Value returns nullable value in case -1 was provided to write NULL into a database
func (s Score) Value() (driver.Value, error) {
	if s == -1 {
		return nil, nil
	}

	return int64(s), nil
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

// Validate validates all filled values regarding main review entity
func (r ReviewUpdate) Validate() error {
	if r.Name != "" {
		if err := validateName(r.Name); err != nil {
			return err
		}
	}

	if r.Email != "" {
		if err := validateEmail(r.Email); err != nil {
			return err
		}
	}

	if r.Content != "" {
		if err := validateContent(r.Content); err != nil {
			return err
		}
	}

	if r.Category != "" {
		if err := validateCategory(r.Category); err != nil {
			return err
		}
	}

	if err := validateScore(Score(r.Score)); err != nil {
		return err
	}
	return nil
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

	if err := validateCategory(r.Category.String); err != nil {
		return err
	}

	if err := validateScore(r.Score); err != nil {
		return err
	}

	return nil
}
