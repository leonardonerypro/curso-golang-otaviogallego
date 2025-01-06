package main

import "fmt"

func main() {
	fmt.Println("Ponteiros - É uma referência a endereço de memória")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro) // sem "*" o ponteiro exibe o endereço de memória da variável referenciada

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // com "*" o ponteiro exibe o valor da variável referenciada
}
