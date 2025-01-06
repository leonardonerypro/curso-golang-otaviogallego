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
}
