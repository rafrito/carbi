package main

import (
	"database/sql"
	"encoding/json"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type Estoque struct {
	Carro string
	Cor   string
	Ano   string
	Preço string
}

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
	m := make([]interface{}, 0, 4)

	// Alguns itens
	m1 := Estoque{"Ferrari-F200", "Vermelha", "2012", "100000.0"}
	m = append(m, m1)

	m1 = Estoque{"Mustang-GT", "Verde", "2012", "80000.0"}
	m = append(m, m1)

	m1 = Estoque{"Golzinho", "Preta", "2015", "30000.0"}
	m = append(m, m1)

	m1 = Estoque{"Fusca", "Preta", "1979", "10000.0"}
	m = append(m, m1)

	db, err := sql.Open("mysql", OrigemDados("carbi"))
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	for _, j := range m {
		err = InsereDado(db, "Estoque", structToMap(j))
		if err != nil {
			panic(err)
		}
	}
}

func structToMap(s interface{}) (m map[string]string) {
	j, _ := json.Marshal(s)
	json.Unmarshal(j, &m)
	return
}
