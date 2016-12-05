package main

import "fmt"

// Persona es una estructura
type Persona struct {
	Nombre string
	Edad   uint8
	Emails []string
}

func main() {
	/*
		    var persona1 Persona
			persona1.Nombre = "Pedro"
			persona1.Edad = 20
			fmt.Println(persona1.Nombre)
	*/

	emails := []string{"a@b.com", "z@b.com"}

	persona2 := Persona{
		"Pablo",
		33,
		emails,
	}
	fmt.Println(persona2)

}
