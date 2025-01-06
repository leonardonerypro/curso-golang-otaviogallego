package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"]) // para acessar uma chave

	usuario2 := map[string]map[string]string{
		"aluno": {
			"nome":      "João",
			"sobrenome": "Pereira",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "Centro",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "curso")
	fmt.Println(usuario2)

	usuario2["filiacao"] = map[string]string{
		"pai": "Antonio",
		"mae": "Maria",
	}
	fmt.Println(usuario2)

	fmt.Println(usuario2["filiacao"])
	fmt.Println(usuario2["filiacao"]["mae"])
}
