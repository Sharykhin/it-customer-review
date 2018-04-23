package entity

import "encoding/json"

type (
	// NullBool detects whether value was provided or not.
	NullBool struct {
		Valid bool
		Value bool
	}
	// NullString provide nullable value
	NullString struct {
		Valid bool
		Value string
	}
)

// UnmarshalJSON implements Unmarshaler interface to allow not to provide value at all
func (n *NullString) UnmarshalJSON(b []byte) error {
	if n == nil {
		n.Valid = false
		n.Value = ""
		return nil
	}
	n.Valid = true
	err := json.Unmarshal(b, &n.Value)
	return err
}

// MarshalJSON returns nullabel value
func (n *NullString) MarshalJSON() ([]byte, error) {
	if n == nil || !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Value)
}

// UnmarshalJSON implements Unmarshaler interface to allow not to provide value at all
func (n *NullBool) UnmarshalJSON(b []byte) error {
	if n == nil {
		n.Value = false
		n.Valid = false
		return nil
	}
	n.Valid = true
	err := json.Unmarshal(b, &n.Value)
	return err
}
