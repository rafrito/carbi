package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func falta() {
	fmt.Println("Ainda não programei essa parte!")
}

type server struct {
	db *sql.DB
}

func badRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"message": "bad request"}`))
}

func statusOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "request succesfull"}`))
}

func (s *server) api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		get(w, r, s.db)
	case "DELETE":
		delete(w, r, s.db)
	case "PUT":
		put(w, r, s.db)
	case "PUSH":
		push(w, r, s.db)
	default:
		badRequest(w)
	}
}

func get(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tab := strings.Split(r.URL.Path, "/")
	switch len(tab) {
	case 3:
		switch tab[1] {
		case "estoque":
			getTudo(w, db, "Estoque")
		case "hist":
			getTudo(w, db, "Histórico")
		default:
			badRequest(w)
		}
	case 4:
		switch tab[1] {
		case "estoque":
			conds := strings.Split(tab[2], "=")
			getCond(w, db, "Estoque", conds)
		case "hist":
			conds := strings.Split(tab[2], "=")
			getCond(w, db, "Histórico", conds)
		}
	default:
		badRequest(w)
	}
}

// Deleta elementos de uma tabela
func delete(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tab := strings.Split(r.URL.Path, "/")
	switch len(tab) {
	case 3:
		err := DeletaCarro(db, tab[1])
		if err != nil {
			badRequest(w)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "item deleted"}`))
		}
	default:
		badRequest(w)
	}
}

// Atualiza elementos de uma tabela
func put(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tab := strings.Split(r.URL.Path, "/")
	switch len(tab) {
	case 5:
		cols := strings.Split(tab[2], ",")
		rows := strings.Split(tab[3], ",")
		err := AtualizaDado(db, tab[1], cols, rows)
		if err != nil {
			badRequest(w)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "item updated"}`))
		}
	default:
		badRequest(w)
	}
}

// Coloca novo elemento numa tabela
func push(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tab := strings.Split(r.URL.Path, "/")
	switch len(tab) {
	case 4:
		m := make(map[string]string)
		cols := strings.Split(tab[1], ",")
		data := strings.Split(tab[2], ",")
		for i, col := range cols {
			m[col] = data[i]
		}
		err := InsereDado(db, "Estoque", m)
		if err != nil {
			badRequest(w)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "item created"}`))
		}
	default:
		badRequest(w)
	}
}

// Imprime no ResponseWritter o texto passado em formato json
func toJSON(s [][]string, w http.ResponseWriter) {
	m := make(map[string]interface{})
	cols := s[0]
	rows := s[1:]
	for _, row := range rows {
		temp := make(map[string]interface{})
		for i, el := range row[1:] {
			temp[cols[i+1]] = el
		}
		m[row[0]] = temp
	}
	js, _ := json.Marshal(m)
	fmt.Fprint(w, string(js))
}

// Devolve todos os valores de uma tabela s
func getTudo(w http.ResponseWriter, db *sql.DB, s string) {
	cols := []string{"*"}
	cond := []string{"none"}
	str, err := GetDados(db, s, cols, cond)
	if err != nil {
		badRequest(w)
		return
	}
	toJSON(str, w)
}

// Devolve os valores restringidos por c numa tabela s
func getCond(w http.ResponseWriter, db *sql.DB, s string, c []string) {
	cols := []string{"*"}
	cond := []string{c[0] + "=" + `'` + c[1] + `'`}
	str, err := GetDados(db, s, cols, cond)
	if err != nil {
		badRequest(w)
		fmt.Println(err)
		return
	}
	toJSON(str, w)
}
