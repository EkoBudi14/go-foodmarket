package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Roles:       user.Roles,
		Address:     user.Address,
		HouseNumber: user.HouseNumber,
		PhoneNumber: user.PhoneNumber,
		Token:       user.Token,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Token: user.Token,
	}
}

func UserToEvent(user *entity.User) *model.UserEvent {
	return &model.UserEvent{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
