package extractors

type VersionExtractor interface {
	GetVersion() string
	GetAppName() string
}
