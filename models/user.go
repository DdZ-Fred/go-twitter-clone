package models

import (
	"time"

	gormtypes "github.com/DdZ-Fred/go-twitter-clone/gorm_types"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id                string               `gorm:"notNull;primaryKey" json:"id"`
	Fname             string               `gorm:"notNull" json:"fname"`
	Lname             string               `gorm:"notNull" json:"lname"`
	Email             string               `gorm:"notNull;unique" json:"email"`
	Username          string               `gorm:"notNull;unique" json:"username"`
	BirthDate         time.Time            `gorm:"notNull" json:"birthDate"`
	Bio               null.String          `json:"bio"`
	Country           string               `gorm:"notNull;size:2" json:"country"`
	PreferredLanguage string               `gorm:"notNull;size:5;default:en-GB" json:"preferredLanguage"`
	Website           null.String          `json:"website"`
	Password          string               `gorm:"notNull"`
	Status            gormtypes.UserStatus `sql:"type:user_status" gorm:"type:user_status;default:pending;notNull" json:"status"`
	ConfirmationCode  null.String          `json:"confirmationCode"`
	CreatedAt         time.Time            `gorm:"notNull" json:"createAt"`
	UpdatedAt         time.Time            `gorm:"notNull" json:"updatedAt"`

	Tweets []Tweet `gorm:"foreignKey:UserId"`
}

func (user User) ToUserSafe() UserSafe {
	return UserSafe{
		Id:                user.Id,
		Fname:             user.Fname,
		Lname:             user.Lname,
		Email:             user.Email,
		Username:          user.Username,
		BirthDate:         user.BirthDate,
		Bio:               user.Bio,
		Country:           user.Country,
		PreferredLanguage: user.PreferredLanguage,
		Website:           user.Website,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
}

type UserSafe struct {
	Id                string      `gorm:"primaryKey" json:"id"`
	Fname             string      `gorm:"notNull" json:"fname"`
	Lname             string      `gorm:"notNull" json:"lname"`
	Email             string      `gorm:"notNull;unique" json:"email"`
	Username          string      `gorm:"notNull;unique" json:"username"`
	BirthDate         time.Time   `gorm:"notNull" json:"birthDate"`
	Bio               null.String `json:"bio"`
	Country           string      `gorm:"notNull;size:2" json:"country"`
	PreferredLanguage string      `gorm:"notNull;size:4" json:"preferredLanguage"`
	Website           null.String `json:"website"`
	CreatedAt         time.Time   `gorm:"notNull" json:"createAt"`
	UpdatedAt         time.Time   `gorm:"notNull" json:"updatedAt"`
}
