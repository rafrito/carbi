package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Pacote com funções úteis para a criação
// e manipulação de banco de dados com mysql

// Hist representa a estrutura do banco de dados Histórico
type Hist struct {
	DataOperação string
	Operação     string
	URL          string
}

// Estoque representa a estrutura do banco de dados Estoque
type Estoque struct {
	Carro string
	Cor   string
	Ano   string
	Preço string
}

// SenhaMYSQL pede a senha do servidor root mysql e
// a armazena numa variavel de ambiente SMYSQL
func SenhaMYSQL() {
	fmt.Println("Digite a senha do servidor mysql root:")
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("SMYSQL", s)
}

// OrigemDados devolve uma string usada para abrir o
// banco de dados carbi
func OrigemDados(s string) string {
	return "root:" + os.Getenv("SMYSQL") + "@/" + s
}

// CriaBanco cria um banco de dados com nome s string
func CriaBanco(s string) {
	dSource := "root:" + os.Getenv("SMYSQL") + "@/"
	db, err := sql.Open("mysql", dSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Exec("drop database if exists carbi")
	db.Exec("create database if not exists " + s)
}

// CriaTabela cria uma nova tabela no banco de dados passado
func CriaTabela(db *sql.DB, nome string) {
	com := "create table if not exists " + nome + " (\n"
	com += "ID integer not null auto_increment,\nPRIMARY KEY (ID)\n)"
	_, err := db.Exec(com)
	if err != nil {
		panic(err)
	}
}

// InsereColuna insere colunas especificadas pelo mapa na tabela tab
func InsereColuna(db *sql.DB, tab string, cols map[string]string) {
	var s string
	for nome, tipod := range cols {
		s = "add " + nome + " " + tipod
		db.Exec("alter table " + tab + "\n" + s)
	}
}

// InsereDado coloca os dados na tabela especificados por um map
func InsereDado(db *sql.DB, tab string, m map[string]string) error {
	cols := make([]string, 0, len(m))
	vals := make([]string, 0, len(m))

	for k, v := range m {
		cols = append(cols, k)
		vals = append(vals, v)
	}

	c := strings.Join(cols, ", ")
	v := strings.Join(vals, `', '`)

	s := "insert into " + tab
	s += `(` + c + `) values('` + v + `')`
	_, err := db.Exec(s)
	return err
}

// DeletaCarro delete um carro identificado pel seu ID
func DeletaCarro(db *sql.DB, ideq string) error {
	_, err := db.Exec(`delete from Estoque where ID='` + ideq + `'`)
	return err
}

// GetDados retorna o valor json associado a um select
// em certa tabela, com determinadas colunas e um condicional
func GetDados(db *sql.DB, tab string, c []string, cond []string) ([][]string, error) {
	s := "select " + strings.Join(c, ", ") + "\n"
	s += "from " + tab

	a := make([][]string, 1)
	if cond[0] != "none" {
		s += "\nwhere " + strings.Join(cond, " and ")
	}

	rows, err := db.Query(s)
	if err != nil {
		return a, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return a, err
	}

	r := make([][]string, 0, 50)
	r = append(r, cols)
	rawResult := make([][]byte, len(cols))

	dest := make([]interface{}, len(cols)) // Um slice interface{} temporário

	for i := range rawResult {
		dest[i] = &rawResult[i] // Usado no rows.Scan para apontar vários valores
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return a, err
		}
		result := make([]string, len(cols))
		for j, raw := range rawResult {
			if raw == nil {
				result[j] = "\\N"
			} else {
				result[j] = string(raw)
			}
		}
		r = append(r, result)
	}
	return r, err
}

// AtualizaDado atualiza uma coluna identificada por id
func AtualizaDado(db *sql.DB, id string, c []string, r []string) error {
	s := "update Estoque\n"
	aux := make([]string, 0, 10)
	for i, col := range c {
		aux = append(aux, col+" = "+`'`+r[i]+`'`)
	}
	s += "set " + strings.Join(aux, ", ") + "\n"
	s += "where ID =" + id

	_, err := db.Exec(s)
	if err != nil {
		return err
	}
	return nil
}

// Registra registra as operações de mudança realizadas no Estoque
func Registra(db *sql.DB, tipo string, u string) error {
	t := time.Now()
	reg := Hist{t.String(), tipo, u}

	m := structToMap(reg)
	err := InsereDado(db, "Histórico", m)

	return err
}

func structToMap(s interface{}) (m map[string]string) {
	j, _ := json.Marshal(s)
	json.Unmarshal(j, &m)
	return
}
