package handler

import (
	"log"
	"net/http"

	"github.com/voyadger01/rest-contacts_golang_01/internal/structs"
)

func (s *Server) updateContact(w http.ResponseWriter, id int, updated structs.Contact) {
	if updated.Name == "" || updated.Phone == "" {
		http.Error(w, "Name and phone are required", http.StatusBadRequest)
		return
	}
	query := `UPDATE contacts SET name = $1, phone = $2 WHERE id = $4`
	res, err := s.DataB.Exec(query, updated.Name, updated.Phone, id)
	if err != nil {
		log.Println("Update database err: ", err)
		http.Error(w, "db err", http.StatusInternalServerError)
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
