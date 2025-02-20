package models

import "Teste/AplicacaoWeb/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConectaComBancoDeDados()

	selectTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosOsProdutos.Scan(&id, &quantidade, &nome, &descricao, &preco)

		p.Id = id
		p.Quantidade = quantidade
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descrico, preco, quantidade) values ($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletarProduto(id string) {
	db := db.ConectaComBancoDeDados()
	deletaNoBanco, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletaNoBanco.Exec(id)
	defer db.Close()
}

func EditarProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoDoBanco, err := db.Query("select * produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	AtualizaProduto, err := db.Prepare("update produtos set nome =$1, descricao = $2,preco=$3, quantidade =$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
