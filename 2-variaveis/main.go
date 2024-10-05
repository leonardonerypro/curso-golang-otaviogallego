package main

import "fmt"

func main() {

	// variaveis
	var variavel1 string = "variável 1"
	variavel2 := "variavel 2"
	imprimir(variavel1)
	imprimir(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string
	)
	imprimir(variavel3)
	imprimir(variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	imprimir(variavel5 + " " + variavel6)

	// constantes
	const constante1 string = "constante 1"
	imprimir(constante1)

	// inversao de valores de variaveis ja existentes
	variavel5, variavel6 = variavel6, variavel5
	imprimir(variavel5 + " " + variavel6)

}

func imprimir(texto string) {
	fmt.Println(texto)
}
