package main

import (
	"github.com/tales-lopes-meli/bootcamp/pkg/social"
)

func main() {
	user := social.User{}
	user.SetName("Tales")
	user.SetSurname("Lopes")
	user.SetEmail("tales.lopes@mercadolivre.com")
	user.SetPassword("senha")
	user.SetAge(22)
	user.ToString()
}
