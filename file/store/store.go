package file

import (
	"fmt"
	"os"
)

const (
	FILE_NAME = "teste.csv"
)

func CreateProductFile(name string) error {
	data := []byte("id,price,quantity\n")
	err := os.WriteFile(name, data, 0644)

	return err
}

func WriteProductFile(id string, price float64, quantity int) error {
	dataConcat := fmt.Sprintf("%s,%.2f,%d\n", id, price, quantity)
	dataToWrite := []byte(dataConcat)
	filePath, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	defer filePath.Close()

	if _, errorFile := filePath.Write(dataToWrite); err != nil {
		return errorFile
	}

	return err
}
