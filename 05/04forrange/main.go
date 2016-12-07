package main

import "fmt"

func main() {

	/*
		numeros := []string{"cero", "uno", "dos", "tres"}

		for _, v := range numeros {
			fmt.Println(v)
		}
	*/
	/*
		nombres := map[string]string{"es": "España", "co": "Colombia", "br": "Brasil"}

		for _, v := range nombres {
			fmt.Println(v)
		}
	*/

	/*
		frase := "Hola mundo"

		for _, letra := range frase {
			fmt.Println(string(letra))
		}
	*/

	for _, letra := range "Curso Go desde cero" {
		fmt.Println(string(letra))
	}
}
