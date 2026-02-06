package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/voyadger01/rest-contacts_golang_01/handler"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatalln("database in env required")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("can't open database")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalln("bad connection")
	}
	s := &handler.Server{DataB: db}
	http.HandleFunc("/contacts", s.ContactsHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Server is running on ", port)
	http.ListenAndServe(port, nil)
}
