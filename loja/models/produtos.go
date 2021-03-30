package models

import (
	"github.com/juliocmalvares/loja/db"
)

type Produto struct {
	id              int
	Nome, Descricao string
	Preco           float64
	Qtd             int
}

func FindAllProducts() []Produto {
	db := db.ConectaBanco()
	_select, err := db.Query("select * from produtos;")
	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}
	for _select.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = _select.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtos = append(produtos, Produto{id, nome, descricao, preco, int(quantidade)})

	}
	defer db.Close()

	return produtos
}

func InsertProduct(p Produto) {
	db := db.ConectaBanco()
	ins, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	ins.Exec(p.Nome, p.Descricao, p.Preco, p.Qtd)
	defer db.Close()
}
