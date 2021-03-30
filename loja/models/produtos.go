package models

import (
	"github.com/juliocmalvares/loja/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Qtd             int
}

func FindAllProducts() []Produto {
	db := db.ConectaBanco()
	_select, err := db.Query("select * from produtos order by id asc")
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

func DeleteProduct(id string) {
	db := db.ConectaBanco()
	del, err := db.Prepare("delete from produtos as p where p.id = $1")
	if err != nil {
		panic(err.Error())
	}
	del.Exec(id)
	defer db.Close()
}

func ProductToPut(id string) Produto {
	db := db.ConectaBanco()
	pr, err := db.Query("select * from produtos as p where p.id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	produto := Produto{}
	for pr.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = pr.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produto.Nome = nome
		produto.Id = id
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Qtd = quantidade
	}
	defer db.Close()
	return produto
}

func PutProduct(p Produto) {
	db := db.ConectaBanco()
	_sql, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id = $5")
	if err != nil {
		panic(err.Error())
	}
	_sql.Exec(p.Nome, p.Descricao, p.Preco, p.Qtd, p.Id)

	defer db.Close()
}
