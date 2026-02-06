package handler

import (
	"database/sql"
	"log"
)

func (s *Server) delete(w http.ResponseWriter, r *http.Request, id int) {
	res, err := db.Exec("DELETE FROM ", dbname, "WHERE ", "id = ", id)
	if err != nil {
		log.Fatalln("bad id to delete")
	}
	rowsDeleted, _ := res.RowsAffected()
	log.Println("Rows deleted: ", rowsDeleted)
}
