package gormtypes

import (
	"database/sql/driver"
	"fmt"
)

const (
	UserStatus_Pending UserStatus = "pending"
	UserStatus_Active  UserStatus = "active"
)

type UserStatus string

func (us *UserStatus) Scan(value string) error {
	fmt.Println("[UserStatus][Scan] value" + value)

	*us = UserStatus(value)
	return nil
}

func (us UserStatus) Value() (driver.Value, error) {
	return string(us), nil
}
