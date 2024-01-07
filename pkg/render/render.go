/*******************************************************************************************************
 * Render.go : fichier qui permet d'afficher les pages HTML
 *******************************************************************************************************/

package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// variables
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	// Vérification : le template est-il déjà en cache ?
	_, inMap := tc[t]
	if !inMap {
		// Si non, on crée le template
		err = createTemplate(t)
		if err != nil {
			fmt.Println("Erreur lors de la création du template :", err)
		}
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
}

func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gohtml",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// ajout du template au cahce
	tc[t] = tmpl
	fmt.Println("Template ajouté au cache")
	return nil
}

func RenderData(w http.ResponseWriter, t string, data interface{}) {
	var tmpl *template.Template
	var err error
	// Vérification : le template est-il déjà en cache ?
	_, inMap := tc[t]
	if !inMap {
		// Si non, on crée le template
		err = createTemplate(t)
		if err != nil {
			fmt.Println("Erreur lors de la création du template :", err)
		}
	} else {
		// Template déjà en cache
		fmt.Println("Template déjà en cache")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, data)
}
