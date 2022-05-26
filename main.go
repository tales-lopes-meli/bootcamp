package main

import (
	"github.com/tales-lopes-meli/bootcamp/pkg/social"
	"github.com/tales-lopes-meli/bootcamp/pkg/store"
)

func main() {
	user := social.User{}
	user.SetName("Tales")
	user.SetSurname("Lopes")
	user.SetEmail("tales.lopes@mercadolivre.com")
	user.SetPassword("senha")
	user.SetAge(22)
	user.ToString()

	userStore := store.User{Name: "Tales", Surname: "Lopes", Email: "tales.lopes@mercadolivre.com"}
	nameProduct := "Tesoura"
	var priceProduct float32 = 10.50
	amountProduct := 7
	userStore.AddProduct(&nameProduct, &priceProduct, &amountProduct)
	userStore.AddProduct(&nameProduct, &priceProduct, &amountProduct)
	userStore.AddProduct(&nameProduct, &priceProduct, &amountProduct)
	userStore.AddProduct(&nameProduct, &priceProduct, &amountProduct)
	userStore.ToString()
	userStore.DeleteProducts()
	userStore.ToString()
	userStore.AddProduct(&nameProduct, &priceProduct, &amountProduct)
	userStore.ToString()
}
