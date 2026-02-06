package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	DataB *sql.DB
}

func (s *Server) ContactsHandler (w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id == "" {
			s.GetAll(w, r)
		} else {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Fatalln("bad format of id")
			}
			s.GetOne(w, r, idInt)
		}
	case http.MethodPost:
		s.Post(w, r)
	case http.MethodPut:
		id := r.URL.Query().Get("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln("bad format of id")
		}
		s.Put(w, r, idInt)
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln("bad format of id")
		}
		s.Delete(w, r, idInt)
	}
}
