package handler

type SignInCredential struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthToken struct {
	Token string `json:"token"`
}
