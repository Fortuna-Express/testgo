package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Obtén el puerto de la variable de entorno, o usa 8080 por defecto.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configura la ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola, Mundo!")
	})

	// Inicia el servidor
	log.Printf("Escuchando en el puerto %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
