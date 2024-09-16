package controller

import "echo-server/domain"

type User struct {
	ID    string
	Name  string
	Email string
}

func FromUserDomainObject(_user *domain.User) *User {
	return &User{
		ID:    _user.ID.String(),
		Name:  _user.Name.String(),
		Email: _user.Email.String(),
	}
}

func FromUserDomainObjectSlice(_users []*domain.User) []*User {
	users := make([]*User, len(_users))
	for i, _user := range _users {
		users[i] = FromUserDomainObject(_user)
	}
	return users
}

func ToUserDomainObject(user *User) *domain.User {
	return &domain.User{
		ID:    domain.UserID(user.ID),
		Name:  domain.UserName(user.Name),
		Email: domain.UserEmail(user.Email),
	}
}

func ToUserDomainObjectSlice(users []*User) []*domain.User {
	_users := make([]*domain.User, len(users))
	for i, user := range users {
		_users[i] = ToUserDomainObject(user)
	}
	return _users
}
