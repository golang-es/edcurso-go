package main

import "fmt"

func main() {
	bodegaOrigen := []string{"php", "javascript", "html", "css", "java", "bases de datos", "git"}
	bodegaDestino := []string{}

	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)

	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	// Se deja en la bodega destino
	go func() {
		for {
			// Se entrega el contenido de la fotocopiadora a la variable libro
			libro := <-fotocopiadora

			fotocopiado <- libro

			// Agregar el libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaOrigen) == len(bodegaDestino) {
				// Cerrar el canal fotocopiado
				close(fotocopiado)
			}
		}
	}()

	fmt.Println("** Listado de libros fotocopiados **")
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}

}
