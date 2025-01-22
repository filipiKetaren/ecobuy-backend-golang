package request

import "ecobuy/entities"

type RegisterRequest struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	MembershipStatus string `json:"member_ship_status"`
}

func (rr RegisterRequest) ToEntities() entities.User {
	return entities.User{
		ID:               rr.ID,
		Name:             rr.Name,
		Email:            rr.Email,
		Password:         rr.Password,
		MembershipStatus: rr.MembershipStatus,
	}
}
