package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NullFloat struct {
	sql.NullFloat64
}

// MarshalJSON to output only string or null (removes "Valid" field)
func (ns NullFloat) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Float64)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON to properly parse JSON into our custom type
func (ns *NullFloat) UnmarshalJSON(data []byte) error {
	var s *float64
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s != nil {
		ns.Float64 = *s
		ns.Valid = true
	} else {
		ns.Valid = false
	}
	return nil
}

// Scan makes NullFloat work with `sqlx` (reads from database)
func (ns *NullFloat) Scan(value interface{}) error {
	if value == nil {
		ns.Float64, ns.Valid = 0.0, false
		return nil
	}
	flt, ok := value.(float64)
	if !ok {
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into NullFloat", value)
	}
	ns.Float64, ns.Valid = flt, true
	return nil
}

// Value makes NullFloat work with database inserts/updates
func (ns NullFloat) Value() (driver.Value, error) {
	if ns.Valid {
		return ns.Float64, nil
	}
	return nil, nil
}
