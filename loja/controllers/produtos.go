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

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	produto := models.ProductToPut(r.URL.Query().Get("id"))
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			panic(err.Error())
		}
		qtd, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			panic(err.Error())
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err.Error())
		}

		p := models.Produto{
			Id:        id,
			Nome:      r.FormValue("nome"),
			Descricao: r.FormValue("descricao"),
			Preco:     preco,
			Qtd:       qtd,
		}

		models.PutProduct(p)
	}
	http.Redirect(w, r, "/", 301)
}
