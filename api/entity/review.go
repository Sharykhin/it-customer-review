package entity

import (
	"encoding/json"
)

type (
	// Review is a basic entity that represent user's review
	Review struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		Content   string     `json:"content"`
		Published bool       `json:"published"`
		Score     Score      `json:"score"`
		Category  NullString `json:"category"`
		CreatedAt string     `json:"created_at"`
		UpdatedAt string     `json:"updated_at"`
	}
	// Score is a specific type that returns nil in case -1 is provided
	Score int64
	// NullString returns nullable value for a client
	//NullString string
)

// MarshalJSON returns null in case -1 is provided
func (s Score) MarshalJSON() ([]byte, error) {
	if s == -1 {
		return json.Marshal(nil)
	}

	return json.Marshal(int64(s))
}
