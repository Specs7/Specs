package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bonjour ! Ceci est un serveur web de base en Go.")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Le serveur est démarré sur le port 8080.")
	http.ListenAndServe(":8080", nil)
}
