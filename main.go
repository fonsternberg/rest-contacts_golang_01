package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/voyadger01/rest-contacts_golang_01/handler"
)

func main() {
	var dbname string
	fmt.Println("Enter database name: ")
	n, err := fmt.Scanln(&dbname)
	if err != nil || n > 1 {
		log.Fatalln("bad database name")
	}
	db, err := sql.Open("mysql", dbname)
	for {
		fmt.Println("Enter method: ")
		handler.Handler(db, dbname)
	}
}