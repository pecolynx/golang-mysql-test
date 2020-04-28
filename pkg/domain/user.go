package domain

import "time"

type User interface {
	ID() uint
	Name() string
	Age() int
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type user struct {
	id        uint
	name      string
	age       int
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(id uint, name string, age int, createdAt time.Time, updatedAt time.Time) User {
	return &user{
		id:        id,
		name:      name,
		age:       age,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *user) ID() uint {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Age() int {
	return u.age
}

func (u *user) CreatedAt() time.Time {
	return u.createdAt
}

func (u *user) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *user) String() string {
	return u.name
}
