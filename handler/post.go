package handler

import (
	"net/http"

	"github.com/voyadger01/rest-contacts_golang_01/structs"
)

func (s *Server) post(w http.ResponseWriter, r *http.Request, newcnt structs.Contact) {
	result, err := s.DataB.Exec("INSERT INTO contacts VALUES", newcnt.Id, newcnt.Name, newcnt.Phone)
	if err != nil {
		http.Error(w, "can't insert values", http.StatusInternalServerError)
	}
	rowsCheck, _ := result.RowsAffected()
	log.Println("Rows added: ", rowsCheck)
}
