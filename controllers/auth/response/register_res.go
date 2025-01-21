package response

import "ecobuy/entities"

type RegisterResponse struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	MembershipStatus string `json:"member_ship_status"`
}

func RegisterFromEntities(user entities.User) RegisterResponse {
	return RegisterResponse{
		ID:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		MembershipStatus: user.MembershipStatus,
	}
}
