package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/voyadger01/rest-contacts_golang_01/structs"
)

type Server struct {
	DataB *sql.DB
}

func (s *Server) ContactsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")

		if id == "" {
			s.getAll(w)
			return
		} else {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "bad id(must be integer)", http.StatusBadRequest)
				return
			}
			s.getOne(w, idInt)
		}

	case http.MethodPost:
		var contact structs.Contact

		err := json.NewDecoder(r.Body).Decode(&contact)
		if err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if contact.Name == "" {
			http.Error(w, "name required", http.StatusBadRequest)
		}
		if contact.Phone == "" {
			http.Error(w, "phone required", http.StatusBadRequest)
		}
		s.createContact(w, contact)
		return

	case http.MethodPut:
		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")
		phone := r.URL.Query().Get("phone")
		idInt, err := strconv.Atoi(id)
		updated := structs.Contact{Id: idInt, Name: name, Phone: phone}

		if err != nil {
			log.Fatalln("bad format of id")
		}
		s.updateContact(w, idInt, updated)

	case http.MethodDelete:
		id := r.URL.Query().Get("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln("bad format of id")
		}

		s.deleteContact(w, idInt)
	}
}
