package main

import "fmt"

type CustomError struct {
	status int
	msg    string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %s with status %d\n", e.msg, e.status)
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

func main() {
	var salario int = 20000
	status, err := salaryTest(salario)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Println(status)
	}
}
