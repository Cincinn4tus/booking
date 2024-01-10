package sql

import "net/http"

/*
**********************************************************************************************************************
Toutes les fonctions de vérification des données des formulaires sont regroupées dans ce fichier. En fonction du nom du formulaire,
la fonction va appeler la fonction de vérification correspondante.
***********************************************************************************************************************/

func FormName(w http.ResponseWriter, r *http.Request) {
	form := r.FormValue("formName")
	switch form {
	case "register":
		VerifyRegister(w, r)
	case "login":
		VerifyLogin(w, r)
	case "new-event":
		VerifyNewEvent(w, r)
	}
}

/***********************************************************************************************************************
* Structures de données pour la gestion des formulaires
***********************************************************************************************************************/

// UserData : structure de données pour la gestion des données du formulaire d'inscription et de connexion
type UserData struct {
	Firstname, Lastname, Email, Password, Message string
}

// EventData : structure de données pour la gestion des données du formulaire de création d'événement
type EventData struct {
	Host, RoomName, EventTitle, EventCategory, EventDescription, EventBeginDate, EventBeginHour, EventEndDate, EventEndHour, Message string
}
