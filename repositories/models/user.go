package models

import (
	"ecobuy/entities"
	"time"
)

type User struct {
	ID               int       `gorm:"primarykey;autoIncrement"`
	Name             string    `gorm:"type:varchar(255);not null"`
	Email            string    `gorm:"type:varchar(255);unique;not null"`
	Password         string    `gorm:"type:varchar(255);not null"`
	MembershipStatus string    `gorm:"type:varchar(255);not null"`
	Points           int       `gorm:"type:int(255);not null"`
	CreateAt         time.Time `gorm:"autoCreateTime"`
	UpdateAt         time.Time `gorm:"autoCreateTime"`
}

func FromEntitiesUser(user entities.User) User {
	return User{
		ID:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		Password:         user.Password,
		MembershipStatus: user.MembershipStatus,
		Points:           user.Points,
		CreateAt:         time.Now(),
		UpdateAt:         time.Now(),
	}
}

func (user User) ToEntities() entities.User {
	return entities.User{
		ID:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		Password:         user.Password,
		MembershipStatus: user.MembershipStatus,
		Points:           user.Points,
		CreateAt:         time.Now(),
		UpdateAt:         time.Now(),
	}
}
