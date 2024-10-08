package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("arquivo structs")

	var u usuario
	u.nome = "Leonardo"
	u.idade = 47
	fmt.Println(u)

	enderecoU2 := endereco{"Rua das Flores", 123}
	u2 := usuario{
		"Leonardo",
		47,
		enderecoU2,
	}
	fmt.Println(u2)

	u3 := usuario{idade: 47, nome: "Leonardo"}
	fmt.Println(u3)
}
