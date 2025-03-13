package models

import "time"

type Context struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateContext struct {
	Name string `json:"name"`
}
