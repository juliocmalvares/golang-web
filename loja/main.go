package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Qtd             int32
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Camiseta Bonita", Preco: 29.99, Qtd: 10},
		{"Tênis", "Confortável", 299.99, 10},
		{"Fone Bluetooth", "Muito bão", 59.99, 2},
	}

	temp.ExecuteTemplate(w, "index", produtos)
}
