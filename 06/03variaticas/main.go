package main

import "fmt"

func main() {
	variatica("Daniel", "Alvaro", "Alexys", "Carol")
}

func variatica(nombres ...string) {
	// slice de string
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
}

func saludar(a int, nombres ...string) {
	// Hacer algo
}
