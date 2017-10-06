package models

type Health struct {
	Uptime string `json:"uptime"`
	DbOk   string `json:"dbOk,omitempty"`
}
