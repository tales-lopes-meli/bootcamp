package social

import (
	"fmt"
)

type User struct {
	name     string
	surname  string
	age      int
	email    string
	password string
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) SetSurname(surname string) {
	user.surname = surname
}

func (user *User) SetAge(age int) {
	user.age = age
}

func (user *User) SetEmail(email string) {
	user.email = email
}

func (user *User) SetPassword(password string) {
	user.password = password
}

func (user *User) ToString() {
	fmt.Printf("Name: %s\nSurname: %s\nEmail: %s\nPassword: %s\nAge: %d\n", user.name, user.surname, user.email, user.password, user.age)
}
