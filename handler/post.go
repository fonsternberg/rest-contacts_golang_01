package handler

import (
	"database/sql"
	"log"

	"github.com/voyadger01/rest-contacts_golang_01/structs"
)

func (s *Server) Post(w http.ResponseWriter, r *http.Request) {
	result, err := db.Exec("INSERT INTO ", dbname, "VALUES", cnt.Name, cnt.Phone)
	if err != nil {
		log.Fatalln("bad post")
	}
	rowsCheck, _ := result.RowsAffected()
	log.Println("Rows added: ", rowsCheck)
}
