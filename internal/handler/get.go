package handler

import (
	"encoding/json"
	"net/http"

	"github.com/voyadger01/rest-contacts_golang_01/internal/structs"
)

func (s *Server) getOne(w http.ResponseWriter, id int) {
	rows, err := s.DataB.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var c structs.Contact
		err := rows.Scan(&c.Id, &c.Name, &c.Phone)
		if err != nil {
			http.Error(w, "scan err", http.StatusInternalServerError)
		}
		if c.Id == id {
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(c)
			if err != nil {
				http.Error(w, "encode error", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "bad id", http.StatusInternalServerError)
}

func (s *Server) getAll(w http.ResponseWriter) {
	rows, err := s.DataB.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var contacts []structs.Contact
	for rows.Next() {
		var c structs.Contact
		err := rows.Scan(&c.Id, &c.Name, &c.Phone)
		if err != nil {
			http.Error(w, "scan error", http.StatusInternalServerError)
		}
		contacts = append(contacts, c)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contacts)
	if err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
	}
}
