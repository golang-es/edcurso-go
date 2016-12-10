package main

import (
	"fmt"

	"github.com/golang-es/edcurso-go/06/07paquetes/despedida"
	"github.com/golang-es/edcurso-go/06/07paquetes/saludar"
)

func main() {
	nombre := "gente del futuro"

	saludar.Saludar(nombre)

	saludar.MeVes = "Esto es un texto asignado desde el main"
	fmt.Println(saludar.MeVes)

	despedida.Despedirse(nombre)
}
