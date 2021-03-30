package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/juliocmalvares/loja/models"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.FindAllProducts()
	temp.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			panic(err.Error())
		}

		qtd, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			panic(err.Error())
		}
		p := models.Produto{
			Nome:      r.FormValue("nome"),
			Descricao: r.FormValue("descricao"),
			Preco:     preco,
			Qtd:       qtd,
		}

		models.InsertProduct(p)
	}
	http.Redirect(w, r, "/", 301)
}
