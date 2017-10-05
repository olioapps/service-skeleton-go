package extractors

type DbDialect string

const (
	Mysql    DbDialect = "mysql"
	Postgres DbDialect = "postgres"
)

type DbExtractor interface {
	ExtractDialect() DbDialect
	ExtractConnectionString() string
}
