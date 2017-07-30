package models

type ErrorEnvelope struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}
