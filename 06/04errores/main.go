package main

import (
	"errors"
	"fmt"
)

func generar(a int) (bool, error) {
	if a < 500 {
		return true, nil
	} else {
		return false, errors.New("No es menor a 500")
	}
}

func main() {
	esMenor, err := generar(500)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("El valor devuelto es:", esMenor)
}

