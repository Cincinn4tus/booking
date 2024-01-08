package sql

import (
	"net/http"
)

func FormName(w http.ResponseWriter, r *http.Request) {
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
