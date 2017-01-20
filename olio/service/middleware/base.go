package middleware

type UserExtractor interface {
	ExtractUser(username string, password string, requestId string) (interface{}, error)
	ExtractUserByUsername(username string, requestId string) (interface{}, error)
}
