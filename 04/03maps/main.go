package main

import "fmt"

func main() {
	/*
		idiomas := make(map[string]string)
		idiomas["es"] = "Español"
		idiomas["en"] = "Inglés"
		idiomas["fr"] = "Francés"
	*/

	/*
		idiomas := map[string]string{
			"es": "Español",
			"en": "Inglés",
			"fr": "Francés",
			"pt": "Portugués",
		}
	*/
	/*

		fmt.Println(idiomas)
		delete(idiomas, "en")
		fmt.Println(idiomas)
	*/

	/*
		if idioma, ok := idiomas["pt"]; ok {
			fmt.Println("La posición pt sí existe y vale", idioma)
		} else {
			fmt.Println("La posición pt NO existe")
		}
	*/

	numeros := map[uint16]string{
		1:    "Uno es un número chiquito",
		2016: "Esto es otro número",
		4:    "Aquí la llave es negativa",
	}

	fmt.Println(numeros)
}
