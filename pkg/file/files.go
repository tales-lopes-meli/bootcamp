package files

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Cliente struct {
	ID        int64
	Nome      string
	Sobrenome string
	RG        string
	Telefone  string
	Endereco  string
}

func LoadFile(name string) (string, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		return "", fmt.Errorf("O arquivo %s não foi encontrado", name)
	}
	return string(data), nil
}

func GenerateId() int64 {
	rand.Seed(time.Now().UnixNano())
	id := rand.Int63()
	return id
}

func CriarCliente(nome string, sobrenome string, rg string, telefone string, endereco string) Cliente {

	id := GenerateId()

	cliente := Cliente{
		ID:        id,
		Nome:      nome,
		Sobrenome: sobrenome,
		RG:        rg,
		Telefone:  telefone,
		Endereco:  endereco,
	}
	return cliente
}

func getDataString(cliente Cliente) string {
	return fmt.Sprintf("%d|%s|%s|%s|%s|%s\n", cliente.ID, cliente.RG, cliente.Nome, cliente.Sobrenome, cliente.Endereco, cliente.Telefone)
}

func verifyClient(fileName string, rg string) (bool, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return false, errors.New("error: file not found")
	} else {
		isClient := false
		stringData := string(data)
		clientesData := strings.Split(stringData, "\n")
		for _, cliente := range clientesData {
			clientData := strings.Split(cliente, "|")
			if len(clientData) != 1 {
				if clientData[1] == rg {
					isClient = true
				}
			}
		}
		return isClient, nil
	}
}

func AddCliente(fileName string, cliente Cliente) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	clienteCadastrado, _ := verifyClient(fileName, cliente.RG)

	if clienteCadastrado {
		fmt.Println("Cliente já está cadastrado no sistema")
		return
	} else {
		dataString := getDataString(cliente)
		if _, err = file.WriteString(dataString); err != nil {
			panic(err)
		}
	}
}
