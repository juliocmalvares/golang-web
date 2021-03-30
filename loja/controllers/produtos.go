package controllers

import (
	"net/http"
	"text/template"

	"github.com/juliocmalvares/loja/models"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.FindAllProducts()
	temp.ExecuteTemplate(w, "index", produtos)
}
