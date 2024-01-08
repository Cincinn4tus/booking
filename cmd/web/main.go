package main

import (
	"booking/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Main : point d'entrée de l'applicationè
func main() {
	// Gestionnaire de fichiers statiques
	staticFileHandler := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileHandler))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/reservation", handlers.Reservation)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/submit", handlers.Submit)

	// Start the server on port 8080
	fmt.Println(fmt.Sprintf("Démarage du serveur sur le port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
