package handlers

import (
	"booking/pkg/render"
	"booking/pkg/sql"
	"fmt"
	"net/http"
	"strconv"
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
	render.RenderTemplate(w, "login.form.gohtml")
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
		http.Error(w, "Erreur lors de la récupération des salles", http.StatusInternalServerError)
		return
	}
	render.RenderData(w, "events.page.gohtml", events)
}

func Register(w http.ResponseWriter, r *http.Request) {
	render.RenderData(w, "register.form.gohtml", nil)
}

// Submit : fonction qui renvoi vers le fichier de traitement du formulaire
func Submit(w http.ResponseWriter, r *http.Request) {
	sql.FormName(w, r)
}

// NewEvent : fonction d'affichage de la page web "Nouvel événement" (avec récupération des données de la base de données)
func NewEvent(w http.ResponseWriter, r *http.Request) {
	rooms, err := sql.GetRooms()
	if err != nil {
		fmt.Print("Erreur lors de la récupération des événements : ", err)
		return
	}
	render.RenderData(w, "new-event.form.gohtml", rooms)
}

// EditEvent : fonction d'affichage de la page web "Modifier un événement" (avec récupération des données de la base de données)
func EditEvent(w http.ResponseWriter, r *http.Request) {
	eventId := r.URL.Query().Get("id")
	// parser eventId en int
	intEventId, err := strconv.Atoi(eventId)

	fmt.Println("Valeur de l'id en int : ", intEventId)
	event, err := sql.ModifyEvent(intEventId)
	if err != nil {
		fmt.Print("Erreur lors de la récupération des événements : ", err)
		return
	}
	render.RenderData(w, "modify-event.form.gohtml", event)
}

// DeleteEvent : fonction de suppression d'un événement
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventId := r.URL.Query().Get("id")
	// parser eventId en int
	intEventId, err := strconv.Atoi(eventId)

	fmt.Println("Valeur de l'id en int : ", intEventId)
	err = sql.DeleteEvent(r, intEventId)
	if err != nil {
		fmt.Print("Erreur lors de la suppression de l'événement : ", err)
		return
	}
	http.Redirect(w, r, "/events", http.StatusSeeOther)
}
