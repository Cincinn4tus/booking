package sql

import (
	"booking/pkg/render"
	"fmt"
	"github.com/gorilla/sessions"
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

	if r.Method == http.MethodPost {
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
			render.RenderData(w, "register.page.gohtml", userData)
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
	}
	// Mettre le message de succès dans la structure UserData
	userData.Message = "Le compte a été créé avec succès ! Vous pouvez vous connecter"
	// Affichage de la page de connexion
	render.RenderData(w, "login.page.gohtml", userData)
	return
}

/**********************************************************************************************************************
 * Fonction VerifyLogin : vérification des données du formulaire de connexion
**********************************************************************************************************************/

func VerifyLogin(w http.ResponseWriter, r *http.Request) {
	// Récupération des données du formulaire
	loginData := LoginData{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
			return
		}
		// Vérification des données de connexion
		db := InitDB()
		stmt, err := db.Prepare("SELECT email, pwd FROM users WHERE email = ? AND pwd = ?")
		if err != nil {
			fmt.Println("Erreur lors de la préparation de la requête :", err)
		}
		defer stmt.Close()

		result, err := stmt.Query(loginData.Email, loginData.Password)
		if err != nil {
			log.Fatal(err)
		}
		defer result.Close()

		// si result est vide, retour à la page de connexion avec un message d'erreur
		if !result.Next() {
			// Mettre le message d'erreur dans la structure LoginData
			loginData.Message = "L'adresse email et / ou le mot de passe est incorrect"
			// Afficher la page de connexion avec le message d'erreur
			render.RenderData(w, "login.page.gohtml", loginData)
			return
		}
		// Si les informations de connexion sont correctes, on crée une session et on redirige vers la page d'accueil
		var sessionCookie = sessions.NewCookieStore([]byte(loginData.Email))

		fmt.Println("Session créée pour l'utilisateur", loginData.Email)
		fmt.Println("Session :", sessionCookie)
	}
}
