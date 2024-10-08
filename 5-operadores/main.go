package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 2 - 1
	var divisao float32 = 4.0 / 3.0
	multiplicacao := 3 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// operacoes com tipos diferentes nao sao permitidos
	var n1 int8 = 4
	var n2 int8 = 6
	var soma2 int8 = n1 + n2
	fmt.Println(soma2)

	// ATRIBUICAO
	var variavel1 string = "string1"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// OPERADORES LOGICOS
	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(falso || verdadeiro)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 20
	numero *= 20
	numero /= 10
	numero %= 4
	fmt.Println(numero)

	// OPERADOR TERNARIO: NAO EXISTE EM GO
	// texto := numero > 5 ? "Maior que 5" : "Menor ou igual 5"
}
