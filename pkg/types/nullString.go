package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NullString struct {
	sql.NullString
}

// MarshalJSON to output only string or null (removes "Valid" field)
func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON to properly parse JSON into our custom type
func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s != nil {
		ns.String = *s
		ns.Valid = true
	} else {
		ns.Valid = false
	}
	return nil
}

// Scan makes NullString work with `sqlx` (reads from database)
func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into NullString", value)
	}
	ns.String, ns.Valid = str, true
	return nil
}

// Value makes NullString work with database inserts/updates
func (ns NullString) Value() (driver.Value, error) {
	if ns.Valid {
		return ns.String, nil
	}
	return nil, nil
}
