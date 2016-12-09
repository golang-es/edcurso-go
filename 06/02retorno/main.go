package main

import "fmt"

func main() {
	/*
	var e int
	e = 5
	f := duplicar(e)
	fmt.Println("El doble de", e, "es", f)
	*/
	nombre, edad := retornar()
	fmt.Println(nombre, edad)
}

/*
func duplicar(a int) int {
	return a * 2
}
*/

/*
func retornar() (string, int) {
	n := "Daniel"
	e := 30
	return n, e
}
*/


func retornar() (n string, e int) {
	e = 30
	n = "Daniel"
	return
}

