package api

type TokenValidator interface {
	IsTokenBlacklisted(string) (bool, error)
}
