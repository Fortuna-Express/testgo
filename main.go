package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Si no se encuentra la variable de entorno PORT, usamos el puerto 8080
	}

	http.HandleFunc("/", handler)
	log.Println("Listening on port", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
