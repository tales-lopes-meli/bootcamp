package main

import (
	"github.com/tales-lopes-meli/bootcamp/pkg/price"
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

	var products []price.Product
	product := price.Product{Name: "Tesoura", Price: 15.5, Amount: 5}
	products = append(products, product)
	product = price.Product{Name: "Borracha", Price: 2.5, Amount: 15}
	products = append(products, product)
	product = price.Product{Name: "Lápis", Price: 4.25, Amount: 55}
	products = append(products, product)

	var services []price.Service
	service := price.Service{Name: "Montagem", Price: 150.0, WorkedMinute: 20}
	services = append(services, service)
	service = price.Service{Name: "Pintura", Price: 230.0, WorkedMinute: 180}
	services = append(services, service)

	var maintenances []price.Maintenance
	maintenance := price.Maintenance{Name: "Jardinagem", Price: 150.0}
	maintenances = append(maintenances, maintenance)
	maintenance = price.Maintenance{Name: "Limpeza", Price: 120.0}
	maintenances = append(maintenances, maintenance)

}
