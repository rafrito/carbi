package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Iniciando servidor carbi...")
	SenhaMYSQL()

	db, err := sql.Open("mysql", OrigemDados("carbi"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server{db: db}
	http.HandleFunc("/", s.api)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
