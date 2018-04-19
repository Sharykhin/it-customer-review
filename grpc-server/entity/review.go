package entity

import "github.com/Sharykhin/it-customer-review/grpc-server/util"

// Review is a basic entity that represent user's review
type Review struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Content   string   `json:"content"`
	Published bool     `json:"published"`
	Score     uint64   `json:"score"`
	Category  string   `json:"category"`
	CreatedAt JSONTime `json:"created_at"`
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
