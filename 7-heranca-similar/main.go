package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança - similar")

	p1 := pessoa{
		"Leonardo",
		"Nery",
		47,
		175,
	}

	e1 := estudante{
		p1,
		"Análise de Sistemas",
		"Unicesumar",
	}

	fmt.Println(e1)
}
