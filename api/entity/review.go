package entity

import (
	"encoding/json"

	"github.com/Sharykhin/it-customer-review/api/util"
)

// Review is a basic entity that represent user's review
type Review struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Content   string     `json:"content"`
	Published bool       `json:"published"`
	Score     uint64     `json:"score"`
	Category  NullString `json:"category"`
	CreatedAt JSONTime   `json:"created_at"`
}

type NullString string

func (t NullString) MarshalJSON() ([]byte, error) {
	if t == "" {
		return json.Marshal(nil)
	}
	return []byte(t), nil
}

// NewReview returns a new instance of Review entity with filled some properties
func NewReview() *Review {
	uuid, err := util.NewUUID()
	if err != nil {
		panic("could not generate uuid")
	}

	return &Review{
		ID:        uuid,
		Published: false,
	}
}
