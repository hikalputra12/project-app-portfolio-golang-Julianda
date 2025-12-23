package model

import "time"

type Message struct {
	ID        int
	FullName  string
	Email     string
	Message   string
	CreatedAt time.Time
}
