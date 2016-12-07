package main

import "fmt"

func main() {
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	*/

	/*
		for i := 4; i >= 0; i-- {
			if i == 3 {
				continue
			}
			fmt.Println(i)
		}
	*/
	/*
		for i := 0; i < 5; i++ {
			if i == 2 {
				fmt.Println("i vale 2 y se rompe el ciclo")
				break
			}
			fmt.Println(i)
		}
	*/

	// var matriz [3][3]
	matriz := [3][3]int{}
	valor := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			valor++
			matriz[i][j] = valor
		}
	}

	for fila := 0; fila < 3; fila++ {
		for columna := 0; columna < 3; columna++ {
			fmt.Println(matriz[fila][columna])
		}
	}

}
