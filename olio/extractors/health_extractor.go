package extractors

import (
	"time"
)

type UptimeExtractor interface {
	GetUptime() time.Duration
}

type DbHealthExtractor interface {
	GetDbExtractor() DbExtractor
}
