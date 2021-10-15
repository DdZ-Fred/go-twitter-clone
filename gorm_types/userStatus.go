package gormtypes

import "database/sql/driver"

type UserStatus string

func (us *UserStatus) Scan(value interface{}) error {
	*us = UserStatus(value.([]byte))
	return nil
}

func (us UserStatus) Value() (driver.Value, error) {
	return string(us), nil
}
