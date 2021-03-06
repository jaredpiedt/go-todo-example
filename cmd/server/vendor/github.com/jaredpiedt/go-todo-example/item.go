package todo

import (
	"time"
)

// Item represents a single todo item.
type Item struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Completed   bool      `json:"completed,omitemtpy"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
