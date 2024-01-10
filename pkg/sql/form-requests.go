package sql

import (
	"booking/pkg/render"
	"fmt"
	"log"
	"net/http"
)

/**********************************************************************************************************************
 * Fonction VerifyRegister : vérification des données du formulaire d'inscription
 *********************************************************************************************************************/

func VerifyRegister(w http.ResponseWriter, r *http.Request) {
	// Récupération des données du formulaire
	userData := UserData{
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
		return
	}

	// L'email est-il déjà utilisé ?
	db := InitDB()
	stmt, err := db.Prepare("SELECT email FROM users WHERE email = ?")
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête :", err)
	}
	defer stmt.Close()

	result, err := stmt.Query(userData.Email)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	// si result n'est pas vide, retour à la page d'inscription avec un message d'erreur
	if result.Next() {
		// Mettre le message d'erreur dans la structure UserData
		userData.Message = "Le compte n'a pas été créé. L'adresse email est déjà utilisée"
		// Afficher la page d'inscription avec le message d'erreur
		render.RenderData(w, "register.form.gohtml", userData)
		return
	}

	// Insertion des données dans la base de données
	stmt, err = db.Prepare("INSERT INTO users (firstname, lastname, email, pwd) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userData.Firstname, userData.Lastname, userData.Email, userData.Password)
	if err != nil {
		log.Fatal(err)
	}

	// Mettre le message de succès dans la structure UserData
	userData.Message = "Le compte a été créé avec succès ! Vous pouvez vous connecter"
	// Affichage de la page de connexion
	render.RenderData(w, "login.form.gohtml", userData)
	return
}

/**********************************************************************************************************************
 * Fonction VerifyLogin : vérification des données du formulaire de connexion
**********************************************************************************************************************/

func VerifyLogin(w http.ResponseWriter, r *http.Request) {
	// Récupération des données du formulaire
	userData := UserData{
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
		Firstname: "",
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
		return
	}
	// L'email est-il déjà utilisé ?
	db := InitDB()
	stmt, err := db.Prepare("SELECT firstname, lastname, email, pwd FROM users WHERE email = ? AND pwd = ?")
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête :", err)
	}
	defer stmt.Close()

	result, err := stmt.Query(userData.Email, userData.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	// si result est vide, retour à la page de connexion avec un message d'erreur
	if !result.Next() {
		// Mettre le message d'erreur dans la structure UserData
		userData.Message = "L'adresse email ou le mot de passe est incorrect"
		// Afficher la page de connexion avec le message d'erreur
		render.RenderData(w, "login.form.gohtml", userData)
		return
	}
	// uSER DATA = concaténation Bienvenue + prénom + nom !
	userData.Message = "Vous êtes connecté !"

	// Récupération de toutes les données de l'utilisateur
	err = result.Scan(&userData.Firstname, &userData.Lastname, &userData.Email, &userData.Password)
	if err != nil {
		log.Fatal(err)
	}
	// Affichage de la page d'accueil avec le message de succès
	render.RenderData(w, "home.page.gohtml", userData)
}

/**********************************************************************************************************************
 * Fonction VerifyNewEvent : vérification des données du formulaire de création d'un nouvel événement
**********************************************************************************************************************/

func VerifyNewEvent(w http.ResponseWriter, r *http.Request) {
	// Récupération des données du formulaire
	eventData := EventData{
		Host:             "Anonyme",
		RoomName:         r.FormValue("roomName"),
		EventTitle:       r.FormValue("title"),
		EventCategory:    r.FormValue("category"),
		EventDescription: r.FormValue("description"),
		EventBeginDate:   r.FormValue("begin_date"),
		EventBeginHour:   r.FormValue("begin_hour"),
		EventEndDate:     r.FormValue("end_date"),
		EventEndHour:     r.FormValue("end_hour"),
		Message:          "",
	}

	// Conversion de la date de début de l'événement au format YYYY-MM-DD
	eventData.EventBeginDate = eventData.EventBeginDate[6:10] + "-" + eventData.EventBeginDate[3:5] + "-" + eventData.EventBeginDate[0:2]

	// Conversion de la date de fin de l'événement au format YYYY-MM-DD
	eventData.EventEndDate = eventData.EventEndDate[6:10] + "-" + eventData.EventEndDate[3:5] + "-" + eventData.EventEndDate[0:2]

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
		return
	}

	// Insertion des données dans la base de données
	db := InitDB()
	stmt, err := db.Prepare("INSERT INTO events (host, room_name, event_title, event_category, event_description, begin_date, end_date, begin_hour, end_hour) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventData.Host, eventData.RoomName, eventData.EventTitle, eventData.EventCategory, eventData.EventDescription, eventData.EventBeginDate, eventData.EventEndDate, eventData.EventBeginHour, eventData.EventEndHour)
	if err != nil {
		log.Fatal(err)
	}

	// Mettre le message de succès dans la structure EventData
	eventData.Message = "L'événement a été créé avec succès !"
	// Affichage de la page d'accueil avec le message de succès
	render.RenderData(w, "home.page.gohtml", eventData)
}

/**********************************************************************************************************************
 * Fonction EditEvent : modification des données d'un événement
**********************************************************************************************************************/

func EditEvent(w http.ResponseWriter, r *http.Request) {

}
