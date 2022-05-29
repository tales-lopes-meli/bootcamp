package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	status int
	msg    string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s with status %d\n", e.msg, e.status)
}

func salaryTest(salario int) (int, error) {
	if salario < 15000 {
		return 500, &CustomError{
			status: 500,
			msg:    fmt.Sprintf("Salário de %d é menor que 15000\n", salario),
		}
	}
	return 200, nil
}

func salaryVerify(salario int) (int, error) {
	if salario < 15000 {
		return 500, errors.New(fmt.Sprintf("Salário de %d é bem menor que 15000 with status 500\n", salario))
	}
	return 200, nil
}

func main() {
	var salario int = 2000
	status, err := salaryTest(salario)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Println(status)
	}

	// EX2
	status, err = salaryVerify(salario)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Println(status)
	}
}
