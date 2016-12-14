package main

import "fmt"

func main() {
	fmt.Println("Conectando a la base de datos...")
	defer cerrarDB()
	fmt.Println("Consultamos información, set de datos")
	defer cerrarSetDeDatos()
}

func cerrarDB() {
	fmt.Println("Cerrar la base de datos")
}

func cerrarSetDeDatos() {
	fmt.Println("Cerra el set de datos")
}
