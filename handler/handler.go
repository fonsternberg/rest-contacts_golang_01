package handler

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
	"net/http"

	"github.com/voyadger01/rest-contacts_golang_01/responces"
)

type Server struct {
	DataB *sql.DB
}

func (s *Server) ContactsHandler (w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":
		
	case "POST":

	case "PUT":

	case "DELETE":

	}
}
