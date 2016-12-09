package main

import "fmt"

func main() {
	consecutivo := secuencia()
	for i := 0; i < 10; i++ {
		fmt.Println(consecutivo())
	}
}

func secuencia() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
