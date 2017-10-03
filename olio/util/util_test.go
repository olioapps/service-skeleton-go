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
		{name: "mysql", connectionString: "root:somepw@/someapp?parseTime=true", expectedResult: "mysql"},
		{name: "no connection string", connectionString: "", expectedResult: "Must have db connection string"},
		{name: "un parseable connection string", connectionString: "bad:connectionstring", expectedResult: "Unable to parse db connection string"},
		{name: "malformed connection string", connectionString: "badconnectionstring", expectedResult: "Malformed db connection string"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dialect, err := DbDialect(tc.connectionString)
			if tc.name == "postgres" || tc.name == "mysql" {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResult, dialect)
				return
			}

			assert.NotNil(t, err)
			assert.Equal(t, tc.expectedResult, err.Error())
		})
	}
}
