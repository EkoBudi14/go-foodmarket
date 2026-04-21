package model

type UserResponse struct {
	ID               string  `json:"id,omitempty"`
	Name             string  `json:"name,omitempty"`
	Email            string  `json:"email,omitempty"`
	Roles            *string `json:"roles,omitempty"`
	Address          *string `json:"address,omitempty"`
	HouseNumber      *string `json:"house_number,omitempty"`
	PhoneNumber      *string `json:"phone_number,omitempty"`
	Token            string  `json:"token,omitempty"`
	ProfilePhotoPath *string `json:"profile_photo_path,omitempty"`
	CreatedAt        int64   `json:"created_at,omitempty"`
	UpdatedAt        int64   `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type UpdateProfilePhotoRequest struct {
	ID string `json:"-" validate:"required,max=100"`
}

type RegisterUserRequest struct {
	ID          string `json:"id" validate:"required,max=100"`
	Password    string `json:"password" validate:"required,max=100"`
	Name        string `json:"name" validate:"required,max=100"`
	Email       string `json:"email" validate:"required,email"`
	Address     string `json:"address" validate:"required,max=100"`
	HouseNumber string `json:"house_number" validate:"required,max=100"`
	PhoneNumber string `json:"phone_number" validate:"required,max=100"`
	City        string `json:"city" validate:"required,max=100"`
}

type UpdateUserRequest struct {
	ID       string `json:"-" validate:"required,max=100"`
	Password string `json:"password,omitempty" validate:"max=100"`
	Name     string `json:"name,omitempty" validate:"max=100"`
}

type LoginUserRequest struct {
	ID       string `json:"id" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}
