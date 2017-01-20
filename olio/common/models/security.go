package models

type Permission struct {
	Type       string
	Operation  string
	ObjectType string
	ObjectID   string
}

type AccessContext struct {
	SystemAccess bool
	UserID       int
	RequestID    string
	Permissions  []*Permission
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Device   string `json:"device"`
}
