package handler

import (
	"net/http"
)

func (s *Server) post(w http.ResponseWriter, r *http.Request) {
	result, err := s.DataB.Exec("INSERT INTO contacts VALUES", cnt.Name, cnt.Phone)
	if err != nil {
		log.Fatalln("bad post")
	}
	rowsCheck, _ := result.RowsAffected()
	log.Println("Rows added: ", rowsCheck)
}
