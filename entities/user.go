package entities

import "time"

type User struct {
	ID               int
	Name             string
	Email            string
	Password         string
	MembershipStatus string
	Points           int
	Token            string
	CreateAt         time.Time
	UpdateAt         time.Time
}
