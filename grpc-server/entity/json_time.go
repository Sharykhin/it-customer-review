package entity

import (
	"encoding/json"
	"time"
)

const jsonTimeFormat = "2006-01-02T15:04:05"

// JSONTime provides nullable time value in a preferable format
type JSONTime time.Time

// MarshalJSON implements Marshaler interface to return nil or time value
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(time.Time(t).UTC().Format(jsonTimeFormat))
}

// IsZero checks whether current date is zero value or not.
func (t JSONTime) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t JSONTime) String() string {
	return time.Time(t).UTC().Format(jsonTimeFormat)
}
