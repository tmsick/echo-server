package handler

import "echo-server/controller"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func FromUserControllerDTO(_user *controller.User) *User {
	return &User{
		ID:    _user.ID,
		Name:  _user.Name,
		Email: _user.Email,
	}
}

func FromUserControllerDTOSlice(_users []*controller.User) []*User {
	users := make([]*User, len(_users))
	for i, _user := range _users {
		users[i] = FromUserControllerDTO(_user)
	}
	return users
}

func ToUserControllerDTO(user *User) *controller.User {
	return &controller.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserControllerDTOSlice(users []*User) []*controller.User {
	_users := make([]*controller.User, len(users))
	for i, user := range users {
		_users[i] = ToUserControllerDTO(user)
	}
	return _users
}
