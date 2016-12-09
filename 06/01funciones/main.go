package main

import "fmt"

func main() {
	var a string
	var e uint8
	a = "Alvaro"
	e = 30
	saludar(a, e)
}

/*
func saludar() {
	// Instrucciones
	fmt.Println("Hola Escuela Digital")
}
*/

/*
func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}
*/

func saludar(nombre string, edad uint8) {
	fmt.Println("Hola", nombre, "Tienes ", edad, "años")
}

