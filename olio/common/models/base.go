package models

import "time"

type OlioModel interface {
	GetID() string
}

type OlioBaseModel struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
