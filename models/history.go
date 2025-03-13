package models

import "time"

type History struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	Role      string    `json:"role"`
	ContextID int       `json:"context_id"`
	CreatedAt time.Time `json:"created_at"`
}

type HistoryEntry struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	Context   Context   `json:"context"`
}

type CreateHistoryEntry struct {
	Message   string `json:"message"`
	Role      string `json:"role"`
	ContextID int    `json:"context_id"`
}
