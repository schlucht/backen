package main

import (
	"net/http"

	"github.com/schlucht/backen/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "home", &templateData{}); err != nil {
		app.errorlog.Println(err)
	}

}

func (app *application) Persons(w http.ResponseWriter, r *http.Request) {	

	pers, err := app.loadAllPersons()
	if err != nil {
		app.errorlog.Println(err)
	}
	data := make(map[string]interface{})
	data["persons"] = pers
	if err := app.renderTemplate(w, r, "persons", &templateData{
		Data: data,
	}); err != nil {
		app.errorlog.Println(err)
	}

}

func (app *application) loadAllPersons()([]*models.Person, error) {
	persons, err := app.DB.GetPersons()
	if err != nil {
		return nil, err
	}
	return persons, nil
}