package models

type OlioUser struct {
	OlioBaseModel
}

func (u *OlioUser) GetID() string {
	return u.ID
}
