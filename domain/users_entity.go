package domain

import "echo-server/repository"

type User struct {
	ID    UserID
	Name  UserName
	Email UserEmail
}

func ToUserRepositoryDTO(user *User) *repository.User {
	return &repository.User{
		ID:    user.ID.String(),
		Name:  user.Name.String(),
		Email: user.Email.String(),
	}
}

func ToUserRepositoryDTOSlice(users []*User) []*repository.User {
	_users := make([]*repository.User, len(users))
	for i, user := range users {
		_users[i] = ToUserRepositoryDTO(user)
	}
	return _users
}

func FromUserRepositoryDTO(_user *repository.User) *User {
	return &User{
		ID:    UserID(_user.ID),
		Name:  UserName(_user.Name),
		Email: UserEmail(_user.Email),
	}
}

func FromUserRepositoryDTOSlice(_users []*repository.User) []*User {
	users := make([]*User, len(_users))
	for i, _user := range _users {
		users[i] = FromUserRepositoryDTO(_user)
	}
	return users
}
