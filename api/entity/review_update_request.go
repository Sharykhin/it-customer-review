package entity

type (
	// ReviewUpdateRequest is an update request with limited number of fields that can be updated
	ReviewUpdateRequest struct {
		Name      NullString `json:"name"`
		Content   NullString `json:"content"`
		Published NullBool   `json:"published"`
	}
)

// Validate validates income request on updating a review
func (rr ReviewUpdateRequest) Validate() error {
	if rr.Name.Valid {
		if err := validateName(rr.Name.Value); err != nil {
			return err
		}
	}

	if rr.Content.Valid {
		if err := validateContent(rr.Content.Value); err != nil {
			return err
		}
	}
	return nil
}
