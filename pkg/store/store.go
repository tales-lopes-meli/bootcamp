package store

import "fmt"

type Product struct {
	name   string
	price  float32
	amount int
}

type User struct {
	Name     string
	Surname  string
	Email    string
	produtos []Product
}

func NewProduct(name *string, price *float32, amount *int) *Product {
	produto := Product{name: *name, price: *price, amount: *amount}
	return &produto
}

func (user *User) AddProduct(name *string, price *float32, amount *int) {
	produto := NewProduct(name, price, amount)
	user.produtos = append(user.produtos, *produto)
}

func (user *User) DeleteProducts() {
	user.produtos = make([]Product, 0)
}

func (user *User) ToString() {
	fmt.Println(user.Name)
	fmt.Println(user.Surname)
	fmt.Println(user.Email)
	fmt.Println("Produtos:")
	for _, produto := range user.produtos {
		fmt.Println(produto)
	}
}
