package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do pack main")
	auxiliar.Func1()

	err := checkmail.ValidateFormat("leonardo#email.com.br")
	if err != nil {
		fmt.Println(err)
	}
}
