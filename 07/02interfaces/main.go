package main

import (
	"github.com/golang-es/edcurso-go/07/02interfaces/animales"
)

func main() {
	var p animales.Perro
	var g animales.Gato
	p.Nombre = "Firulais"
	g.Nombre = "Manchas"
	// AdoptarPerro(p)
	// AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}

/*
func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}

func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}
*/

func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}
