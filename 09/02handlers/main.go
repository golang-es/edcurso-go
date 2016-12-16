package main

import (
	"net/http"
	"fmt"
	"log"
)

type messageHandler struct {
	message string
}

func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	mux := http.NewServeMux()
	m1 := &messageHandler{message: "<h1>Hola desde un handler</h1>"}
	m2 := &messageHandler{message: "<h1>Estos son Perritos</h1>"}
	m3 := &messageHandler{message: "<h1>Estos son gatitos</h1>"}

	mux.Handle("/saludar", m1)
	mux.Handle("/perros", m2)
	mux.Handle("/gatos", m3)

	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
