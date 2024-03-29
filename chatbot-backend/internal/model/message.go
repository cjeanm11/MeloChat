package models

// Conversation struct with Messages slice
type Conversation struct {
	Messages []Message `json:"messages"`
}

// NewConversation constructor function
func NewConversation() Conversation {
	return Conversation{
		Messages: []Message{},
	}
}

// Message struct with user and text fields
type Message struct {
	User bool   `json:"user"`
	Text string `json:"text"`
}
