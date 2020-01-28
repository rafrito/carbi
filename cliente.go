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

func (s *server) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		get(w, r, s.db)
	case "DELETE":
		delete(w, r, s.db)
	default:
		badRequest(w)
	}

}

func get(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tab := strings.Split(r.URL.Path, "/")
	fmt.Println(r.URL.Path)
	fmt.Println(tab)
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

		case "hist":
			// printa historico limitado
		}
	default:
		badRequest(w)
	}
	falta()
}

func delete(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tab := strings.Split(r.URL.Path, "/")
	switch len(tab) {
	case 3:
		err := DeletaCarro(db, tab[1])
		if err != nil {
			badRequest(w)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "item deleted"}`))
		}
	default:
		badRequest(w)
	}
}

func put(w http.ResponseWriter, r *http.Request) {
	tab := strings.Split(r.URL.Path, "/")
	switch len(tab) {
	case 5:
		// altera elemento do estoque
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "bad request"}`))
	}
	falta()
}

func push(w http.ResponseWriter, r *http.Request) {
	tab := strings.Split(r.URL.Path, "/")
	switch len(tab) {
	case 4:
		// cria elemento do estoque
	default:
		badRequest(w)
	}
	falta()
}

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
	fmt.Println(string(js))
	fmt.Fprint(w, string(js))
}

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

func getCond(w http.ResponseWriter, db *sql.DB, s string, c []string) {
	cols := []string{"*"}
	cond := []string{c[0] + "=" + `'` + c[1] + `'`}
	str, err := GetDados(db, s, cols, cond)
	if err != nil {
		badRequest(w)
		return
	}
	toJSON(str, w)
}
