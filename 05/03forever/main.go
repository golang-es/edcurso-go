package main

import "fmt"

func main() {
	// Es infinito, no tiene una condición que lo detenga.
	// Se usa para procesos en los que se requiere escuchar permanentemente
	// como sockets, conexiones, etc.
	for {
		fmt.Println("Hola")
	}
}
