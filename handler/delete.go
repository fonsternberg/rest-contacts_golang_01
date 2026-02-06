package handler

import (
	"net/http"
	"log"
)

func (s *Server) deleteContact(w http.ResponseWriter, id int) {
	res, err := s.DataB.Exec("DELETE FROM contacts WHERE id = ?	", id)
	if err != nil {
		log.Println("delete err: ", err)
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		log.Println("delete err: ", err)
		http.Error(w, "failed to affect rows", http.StatusInternalServerError)
	}
	if rowsDeleted == 0 {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
