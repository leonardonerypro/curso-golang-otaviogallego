package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Array
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{
		"Posição 1",
		"Posição 2",
		"Posição 3",
		"Posição 4",
		"Posição 5",
	}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slice
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 20)
	fmt.Println(slice)

	slice2 := array2[0:1]
	fmt.Println(slice2)

	array2[0] = "Posição alterada"
	fmt.Println(slice2)

	// Arreys Internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
	fmt.Println(slice3)

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
