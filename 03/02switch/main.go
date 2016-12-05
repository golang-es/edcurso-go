package main

import "fmt"

func main() {
	// var a uint8
	// a = 36

	/*
		switch a {
		case 1:
			fmt.Println("Lunes")
		case 2:
			fmt.Println("Martes")
		case 3:
			fmt.Println("Miercoles")
		case 4:
			fmt.Println("Jueves")
		case 5:
			fmt.Println("Viernes")
		case 6:
			fmt.Println("Sábado")
		case 7:
			fmt.Println("Domingo")
		default:
			fmt.Println("No es un día de de la semana")
		}
	*/
	/*
		switch a {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			fmt.Println("Estás entre semana")
		case 6:
			fallthrough
		case 7:
			fmt.Println("Estás en fin de semana :)")
		default:
			fmt.Println("No es un día válido")
		}
	*/
	switch a := 36; a {
	case 5:
		fmt.Println("Estás entre semana")
	default:
		fmt.Println("No es un día válido")
	}
}
