package entity

// ReviewRequest represents income body when a new review is going to be created
type ReviewRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

// Validate validates the current struct
func (rr ReviewRequest) Validate() error {

	if err := validateName(rr.Name); err != nil {
		return err
	}

	if err := validateEmail(rr.Email); err != nil {
		return err
	}

	if err := validateContent(rr.Content); err != nil {
		return err
	}

	return nil
}
