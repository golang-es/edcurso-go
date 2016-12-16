package main

import (
	"net/http"
	"log"
)

func main() {
	/*
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", nil))
	*/
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
