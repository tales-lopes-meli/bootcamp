package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func GetProductInfo(name string) {
	data, err := os.ReadFile(FILE_NAME)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		strData := string(data)
		rowData := strings.Split(strData, "\n")
		totalValue := 0.0
		for idx, row := range rowData {
			if idx != 0 {
				splitData := strings.Split(row, ",")
				if len(splitData) != 1 {
					data, err := strconv.ParseFloat(splitData[1], 64)
					if err == nil {
						totalValue += data
					}
				}
			}
		}
		fmt.Printf("%s", strings.Replace(strData, ",", "\t\t\t", -1))
		fmt.Printf("\t\t\t%.2f\n", totalValue)
	}
}
