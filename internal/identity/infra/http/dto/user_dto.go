package dto

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type ActivateUserRequest struct {
	UserID uint64 `json:"user_id"`
	Token  string `json:"token"`
}

type FindUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SendConfirmationEmailMessage struct {
	UserID uint64 `json:"user_id"`
}
