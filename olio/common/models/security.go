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

func (self *AccessToken) GetID() string {
	return self.ID
}

func (self *AccessToken) SetID(id string) {
	self.ID = id
}

type UserAuth struct {
	Username string `json:"username" jsonapi:"attr,username"`
	Password string `json:"password" jsonapi:"attr,password"`
	Device   string `json:"device" jsonapi:"attr,email"`
}
