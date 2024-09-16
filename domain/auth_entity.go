package domain

type SignInCredential struct {
	Email    UserEmail
	Password UserPassword
}

type AuthToken struct {
	Token string
}
