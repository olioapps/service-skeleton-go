package models

type Permission struct {
	Type      string
	Operation string
	OlioType  OlioType
}

type AccessContext struct {
	SystemAccess bool
	User         *OlioUser
	RequestID    string
	Permissions  []*Permission
}
