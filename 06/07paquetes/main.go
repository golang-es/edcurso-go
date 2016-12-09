package main

import (
	des "github.com/golang-es/edcurso-go/06/07paquetes/despedirse"
	"github.com/golang-es/edcurso-go/06/07paquetes/saludar"
)

func main() {
	var a string
	a = "Estudiante de ED"
	saludar.Saludar(a)
	des.Despedirse()
}
