package handlers

import (
	"booking/pkg/render"
	"booking/pkg/sql"
	"net/http"
)

// Home : fonction d'affichage de la page web "Accueil"
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

// About : fonction d'affichage de la page web "A propos"
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}

// Login : fonction d'affichage de la page web "Connexion"
func Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.gohtml")
}

// Reservation : fonction d'affichage de la page web "Réservations" (avec récupération des données de la base de données)
func Reservation(w http.ResponseWriter, r *http.Request) {
	rooms, err := sql.GetRooms()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des salles", http.StatusInternalServerError)
		return
	}
	render.RenderData(w, "reservation.page.gohtml", rooms)
}

// Events : fonction d'affichage de la page web "Evénements" (avec récupération des données de la base de données)
func Events(w http.ResponseWriter, r *http.Request) {
	events, err := sql.GetEvents()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des événements", http.StatusInternalServerError)
		return
	}
	render.RenderData(w, "events.page.gohtml", events)
}
