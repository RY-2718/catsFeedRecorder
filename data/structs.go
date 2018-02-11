package data

import "time"

type (
	InnerFood struct {
		ID        int
		Name      string
		CreatedAt time.Time
	}

	OuterFood struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
	}
)
