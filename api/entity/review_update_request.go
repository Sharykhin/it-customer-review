package entity

import "encoding/json"

type (
	ReviewUpdateRequest struct {
		Name      string   `json:"name"`
		Email     string   `json:"email"`
		Published NullBool `json:"published"`
	}

	NullBool struct {
		Valid bool
		Value bool
	}
)

func (n *NullBool) UnmarshalJSON(b []byte) error {
	if n == nil {
		n.Value = false
		n.Valid = false
		return nil
	}
	n.Valid = true
	err := json.Unmarshal(b, &n.Value)
	if err != nil {
		return err
	}
	return nil
}
