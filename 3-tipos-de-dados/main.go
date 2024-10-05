package main

import (
	"errors"
	"fmt"
)

func imprimir(dado any) {
	fmt.Println(dado)
}

func main() {

	// INTEIROS
	// int8, int16, int32, int64
	// int => segue arquitetura de bits da maquina
	var numero int64 = -10000000000
	imprimir(numero)

	// uint => unsigned => sem sinal
	var numero2 uint32 = 1000
	imprimir(numero2)

	// alias
	// int32 => rune
	var numero3 rune = 123456
	imprimir(numero3)

	// byte = uint8
	var numero4 byte = 123
	imprimir(numero4)

	// REAL
	// float32, float64 (nao aceita declarar como apenas float)
	var numeroReal1 float32 = 123.45
	imprimir(numeroReal1)
	var numeroReal2 float64 = 1230000000.50
	imprimir(numeroReal2)
	numeroReal3 := 12345.67
	imprimir(numeroReal3)

	// STRING
	// usar sempre aspas duplas
	// aspas simples retorna posicao do caracter na tabela ascii
	// no Go nao existe CHAR
	var str string = "texto1"
	imprimir(str)
	str2 := "texto2"
	imprimir(str2)

	// VALOR ZERO (INICIALIZACAO)
	// string = ""
	// numero = 0
	// bool   = false
	// error  = <nil>

	// BOOLEAN
	var booleano bool = true
	imprimir(booleano)
	booleano = !true
	imprimir(booleano)

	// ERRO (EM GO EH UM TIPO)
	var erro error
	imprimir(erro)
	erro = errors.New("erro interno")
	imprimir(erro)

}
