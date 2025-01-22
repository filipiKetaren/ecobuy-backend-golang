package request

import "ecobuy/entities"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rr LoginRequest) ToEntities() entities.User {
	return entities.User{
		Email:    rr.Email,
		Password: rr.Password,
	}
}
