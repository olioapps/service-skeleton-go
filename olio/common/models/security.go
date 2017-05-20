package models

import "time"

type Permission struct {
	Type       string
	Operation  string
	ObjectType string
	ObjectID   string
}

type AccessContext struct {
	SystemAccess bool
	UserID       string
	RequestID    string
	Permissions  []*Permission
}

type AccessToken struct {
	ID             string     `json:"id" jsonapi:"primary,accessTokens" gorm:"primary_key"`
	Token          string     `json:"token" jsonapi:"attr,token"`
	ExpirationDate *time.Time `json:"expirationDate" jsonapi:"attr,expirationDate"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Device   string `json:"device"`
}
