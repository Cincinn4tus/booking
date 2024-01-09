package sql

import (
	"booking/pkg/render"
	"fmt"
	"log"
	"net/http"
)

func FormName(w http.ResponseWriter, r *http.Request) {
	// vérification de la méthode HTTP
	if r.Method != http.MethodPost {
		http.Error(w, "Mauvaise méthode HTTP", http.StatusMethodNotAllowed)
		return
	}
	form := r.FormValue("formName")

	switch form {
	case "register":
		VerifyRegister(w, r)
	case "login":
		VerifyLogin(w, r)
	}
}

/***********************************************************************************************************************
* Structures de données pour la gestion des formulaires
***********************************************************************************************************************/

// UserData : structure de données pour la gestion des données du formulaire d'inscription
type UserData struct {
	Firstname, Lastname, Email, Password, Message string
}

// LoginData : structure de données pour la gestion des données du formulaire de connexion
type LoginData struct {
	Email, Password, Message string
}

func VerifyLogin(w http.ResponseWriter, r *http.Request) {
	// Récupération des données du formulaire
	userData := UserData{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
		return
	}

	// L'email est-il déjà utilisé ?
	db := InitDB()
	stmt, err := db.Prepare("SELECT * FROM users WHERE email = ?")
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
	if !result.Next() {
		// Mettre le message d'erreur dans la structure UserData
		userData.Message = "Email ou mot de passe incorrect"
		// Afficher la page de connexion avec le message d'erreur
		render.RenderData(w, "login.page.gohtml", userData)
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
	render.RenderData(w, "login.page.gohtml", userData)
	return
}
