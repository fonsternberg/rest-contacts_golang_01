package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/voyadger01/rest-contacts_golang_01/structs"
)

func (s *Server) getOne(w http.ResponseWriter, r *http.Request, id int) {
	rows, err := db.Query("SELECT id,from ")
	if err != nil {
		log.Fatalln("bad query")		
	}
	defer rows.Close()
	contacts := []contact.Contact{}
	for rows.Next(){
		cnt := contact.Contact{}
		err := rows.Scan(&cnt.Name, &cnt.Phone)
		if err != nil {
			log.Println("bad name or phone")
			continue
		}
		contacts = append(contacts, cnt)
	}
	for i := range contacts {
		fmt.Println("Contact ", id, ": ", contacts[i].Name, " ", contacts[i].Phone)
	}
}

func (s *Server) getAll(w http.ResponseWriter, r *http.Request) {
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
