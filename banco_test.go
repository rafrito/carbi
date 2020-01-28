package main

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCriaBanco(t *testing.T) {
	os.Setenv("SMYSQL", "rafs1793")
	CriaBanco("carbi")
	db, err := sql.Open("mysql", OrigemDados("carbi"))
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	if db.Ping() != nil {
		t.Error("Nome de origem database inválido.")
	}
}

func TestCriaTabela(t *testing.T) {
	db, err := sql.Open("mysql", OrigemDados("carbi"))
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	CriaTabela(db, "Estoque")
	CriaTabela(db, "Histórico")
}

func TestInsereColuna(t *testing.T) {
	db, err := sql.Open("mysql", OrigemDados("carbi"))
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	d1 := map[string]string{
		"Carro": "varchar(255)",
		"Cor":   "varchar(255)",
		"Ano":   "int",
		"Preço": "float",
	}

	d2 := map[string]string{
		"Carro":        "varchar(255)",
		"Cor":          "varchar(255)",
		"DataOperacao": "varchar(255)",
		"Operacao":     "varchar(255)",
		"Ano":          "int",
		"IDCarro":      "int",
		"Preço":        "float",
	}
	InsereColuna(db, "Estoque", d1)
	InsereColuna(db, "Historico", d2)
}

func TestInsereDado(t *testing.T) {
	m := make([]map[string]string, 0, 4)
	m1 := map[string]string{
		"Carro": "Ferrari F-200",
		"Cor":   "Vermelha",
		"Ano":   "2001",
		"Preço": "200000.0",
	}
	m = append(m, m1)
	m1 = map[string]string{
		"Carro": "Mustang GT",
		"Cor":   "Preto",
		"Ano":   "2012",
		"Preço": "120000.0",
	}
	m = append(m, m1)
	m1 = map[string]string{
		"Carro": "Golzinho",
		"Cor":   "Preto",
		"Ano":   "2015",
		"Preço": "30000.0",
	}
	m = append(m, m1)
	m1 = map[string]string{
		"Carro": "Fusca",
		"Cor":   "Azul",
		"Ano":   "1980",
		"Preço": "10000.0",
	}
	m = append(m, m1)

	db, err := sql.Open("mysql", OrigemDados("carbi"))
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	for _, j := range m {
		err = InsereDado(db, "Estoque", j)
		if err != nil {
			panic(err)
		}
	}
}
