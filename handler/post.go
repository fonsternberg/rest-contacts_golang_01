package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/voyadger01/rest-contacts_golang_01/structs"
)

func (s *Server) post(w http.ResponseWriter, r *http.Request, contact structs.Contact) {
	result, err := s.DataB.Exec("INSERT INTO contacts (name, phone, email) VALUES (?, ?)", contact.Name, contact.Phone)
	if err != nil {
		log.Println("DB insert error: ", err)
		http.Error(w, "can't insert", http.StatusInternalServerError)
	}
	tempId, _ := result.LastInsertId()
	contact.Id = int(tempId)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
	}
}
