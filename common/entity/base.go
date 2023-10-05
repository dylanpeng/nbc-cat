package entity

import (
	"database/sql/driver"
	"encoding/json"
)

type IEntity interface {
	TableName() string
	PrimaryPairs() []interface{}
	PrimarySeted() bool
	String() string
}

type Int64s []int64

func (c Int64s) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Int64s) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type Strings []string

func (c Strings) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Strings) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
