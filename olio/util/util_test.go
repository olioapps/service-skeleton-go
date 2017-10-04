package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbDialect(t *testing.T) {
	tt := []struct {
		name             string
		connectionString string
		expectedResult   string
	}{
		{name: "postgres", connectionString: "postgres://someuser@localhost/someapp?sslmode=disable", expectedResult: "postgres"},
		{name: "mysql", connectionString: "root:somepw/someapp?parseTime=true", expectedResult: "mysql"},
		{name: "defaults to mysql", connectionString: "hmm", expectedResult: "mysql"},
		{name: "no connection string", connectionString: "", expectedResult: "Must have db connection string"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dialect, err := DbDialect(tc.connectionString)
			if tc.name == "no connection string" {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedResult, err.Error())
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tc.expectedResult, dialect)
		})
	}
}
