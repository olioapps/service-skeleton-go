package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type StringMap map[string]string

func (s StringMap) Value() (driver.Value, error) {
	j, err := json.Marshal(s)
	return j, err
}

func (s *StringMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	i := make(StringMap)
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*s = i

	return nil
}
