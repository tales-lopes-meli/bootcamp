package main

import (
	files "github.com/tales-lopes-meli/bootcamp/pkg/file"
)

func main() {
	// data, err := files.LoadFile("teste1.txt")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(data)
	// }

	cliente := files.CriarCliente("Tales", "Lopes", "1234567890", "13998808312", "Av Pedro Américo, 321")

	files.AddCliente("teste.txt", cliente)
}
