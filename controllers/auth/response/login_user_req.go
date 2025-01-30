package response

import "ecobuy/entities"

type LoginResponse struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	MembershipStatus string `json:"member_ship_status"`
	Token            string `json:"token"`
}

func LoginFromEntities(user entities.User) LoginResponse {
	return LoginResponse{
		ID:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		MembershipStatus: user.MembershipStatus,
		Token:            user.Token,
	}
}
