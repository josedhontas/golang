package entities

import (
	"github.com/google/uuid"
)

type tweet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func NewTweet() *tweet {
	tweet := tweet{
		ID: uuid.New().String(),
	}

	return &tweet
}
