package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id        string      `gorm:"notNull;primaryKey" json:"id"`
	Fname     string      `gorm:"notNull" json:"fname"`
	Lname     string      `gorm:"notNull" json:"lname"`
	Email     string      `gorm:"notNull;unique" json:"email"`
	Username  string      `gorm:"notNull;unique" json:"username"`
	BirthDate time.Time   `gorm:"notNull" json:"birthDate"`
	Bio       null.String `json:"bio"`
	Country   string      `gorm:"notNull;size:2" json:"country"`
	Website   null.String `json:"website"`
	Password  string      `gorm:"notNull"`
	CreatedAt null.Time   `gorm:"notNull" json:"createAt"`
	UpdatedAt null.Time   `gorm:"notNull" json:"updatedAt"`
}

type UserSafe struct {
	Id        string      `gorm:"primaryKey" json:"id"`
	Fname     string      `gorm:"notNull" json:"fname"`
	Lname     string      `gorm:"notNull" json:"lname"`
	Email     string      `gorm:"notNull;unique" json:"email"`
	Username  string      `gorm:"notNull;unique" json:"username"`
	BirthDate null.Time   `gorm:"notNull" json:"birthDate"`
	Bio       null.String `json:"bio"`
	Country   string      `gorm:"notNull;size:2" json:"country"`
	Website   null.String `json:"website"`
	CreatedAt null.Time   `gorm:"notNull" json:"createAt"`
	UpdatedAt null.Time   `gorm:"notNull" json:"updatedAt"`
}
