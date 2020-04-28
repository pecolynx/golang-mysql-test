package domain

import "fmt"

type UserAddParameter struct {
	Name string
	Age  int
}

func NewUserAddParameter(name string, age int) *UserAddParameter {
	return &UserAddParameter{
		Name: name,
		Age:  age,
	}
}

type UserUpdateParameter struct {
	Name string
	Age  int
}

func NewUserUpdateParameter(name string, age int) *UserUpdateParameter {
	return &UserUpdateParameter{
		Name: name,
		Age:  age,
	}
}

type UserNotFoundError struct {
	id   uint
	text string
}

func NewUserNotFoundError(id uint) *UserNotFoundError {
	return &UserNotFoundError{
		id:   id,
		text: fmt.Sprintf("User not found. user ID: %d", id),
	}
}

func (e *UserNotFoundError) Error() string {
	return e.text
}

type UserRepository interface {
	AddUser(user *UserAddParameter) (uint, error)
	UpdateUser(id uint, user *UserUpdateParameter) (bool, error)
	RemoveUser(id uint) (bool, error)
	FindUserByID(id uint) (User, error)
	FindUsers(limit, offset int) ([]User, error)
	FindUnderageUsers(limit, offset int) ([]User, error)
}
