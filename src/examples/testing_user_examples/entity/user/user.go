package user

import "github.com/google/uuid"

type User struct {
	id        uuid.UUID
	name      string
	email     string
	password  string
	birthDate string
}

func (instance *User) ID() uuid.UUID {
	return instance.id
}

func (instance *User) Name() string {
	return instance.name
}

func (instance *User) Email() string {
	return instance.email
}

func (instance *User) Password() string {
	return instance.password
}

func (instance *User) BirthDate() string {
	return instance.birthDate
}
