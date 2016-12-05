package main

import "fmt"

func main() {
	// Slices
	// 1. Apuntador a un array
	// 2. Tamaño (no es fijo)
	// 3. Capacidad
	// var nombres []string
	// make (tipo de dato del slice, tamaño inicial, capacidad inicial)
	// nombres := make([]string, 0)
	nombres := []string{
		"Alvaro",
		"Daniel",
		"Alexys",
	}

	/*
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Daniel")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Alvaro")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Alexys")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Pedro")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Juan")
		fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))

		nombres[3] = "Juan"
	*/
	fmt.Println(nombres)
}
