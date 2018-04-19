package entity

import (
	"time"
)

const jsonTimeFormat = "2006-01-02T15:04:05"

type (
	// JSONTime provides nullable time value in a preferable format
	JSONTime time.Time
)

//String returns datetime in a preferable format
func (t JSONTime) String() string {
	return time.Time(t).UTC().Format(jsonTimeFormat)
}
