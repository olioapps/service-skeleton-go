package models

type Health struct {
	Uptime string `json:"uptime"`
	DbOk   bool   `json:"dbOk,omitempty`
}
