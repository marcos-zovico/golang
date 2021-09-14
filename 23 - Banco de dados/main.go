package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	stringConexao := "root:Password123-@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, error := sql.Open("mysql", stringConexao)

	if error != nil {
		log.Fatalln(error)
	}

	defer db.Close()

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}

	fmt.Println("conectado")
}
