package saludar

import "fmt"

// MeVes es para utilizar desde otro paquete
var MeVes string

var noMeVes string

// Saludar saluda a una persona
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func noVisible() {
	fmt.Println("No me ven desde otro paquete")
}
