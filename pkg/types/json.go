package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONArr []string

func (j *JSONArr) Scan(val interface{}) error {
	if val == nil {
		*j = nil
		return nil
	}
	bytes, ok := val.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONArr value: %v", val)
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONArr) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
