package api

type TokenValidor interface {
	IsTokenBlacklisted(string) (bool, error)
}
