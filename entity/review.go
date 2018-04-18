package entity

// Review is a basic entity that represent user's review
type Review struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Content   string   `json:"content"`
	Published bool     `json:"published"`
	Score     uint64   `json:"score"`
	CreatedAt JSONTime `json:"created_at"`
}
