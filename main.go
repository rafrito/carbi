package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Iniciando servidor carbi...")
	// SenhaMYSQL()
	os.Setenv("SMYSQL", "rafs1793")

	db, err := sql.Open("mysql", OrigemDados("carbi"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server{db: db}
	http.HandleFunc("/", s.home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
