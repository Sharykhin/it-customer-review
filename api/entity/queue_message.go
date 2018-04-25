package entity

import "encoding/json"

type (
	// QueueMessage represent a general message struct
	QueueMessage struct {
		Action  string          `json:"action"`
		Payload json.RawMessage `json:"payload"`
	}
	// AnalyzeMessage is a specific struct that would be used as a Payload of QueueMessage.
	// This struct contains ID of a new review and content that should be analyzed
	AnalyzeMessage struct {
		ID      string `json:"id"`
		Content string `json:"content"`
	}
)
