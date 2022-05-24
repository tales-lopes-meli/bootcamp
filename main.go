package main

import file "github.com/tales-lopes-meli/bootcamp/file/store"

func main() {
	file.CreateProductFile(file.FILE_NAME)
	file.WriteProductFile("1000", 15.5, 123)
	file.WriteProductFile("1001", 5.58, 1263)
	file.WriteProductFile("1002", 132.5, 1523)
	file.WriteProductFile("1003", 215.5, 5123)

	file.GetProductInfo(file.FILE_NAME)
}
