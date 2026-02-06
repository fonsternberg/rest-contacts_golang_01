package handler

import (
	"net/http"
	"log"
)

func (s *Server) delete(w http.ResponseWriter, r *http.Request, id int) {
	res, err := s.DataB.Exec("DELETE FROM contacts WHERE id = ", id)
	if err != nil {
		log.Println("bad id to delete")
		
	}
	rowsDeleted, _ := res.RowsAffected()
	log.Println("Rows deleted: ", rowsDeleted)
}
