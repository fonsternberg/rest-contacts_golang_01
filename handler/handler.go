package handler

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/voyadger01/rest-contacts_golang_01/contact"
	"github.com/voyadger01/rest-contacts_golang_01/responces"
)

func Handler(db *sql.DB, dbname string) {
	var method, id, name, phone string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())

	switch len(parts) {
	case 1:
		method = parts[0]
	case 2:
		method = parts[0]
		id = parts[1]
	case 4:
		method = parts[0]
		id = parts[1]
		name = parts[2]
		phone = parts[3]
	default:
		log.Fatalln("Wrong format of response")	
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalln("Id should be integer")
	}
	newcnt := contact.Contact{Name: name, Id: idInt, Phone: phone}
	switch method {
	case "GET":
		if len(parts) == 2 {
			responces.GetOne(db, dbname, idInt)
		} else if len(parts) == 1 {
			responces.GetAll(db, dbname)
		}
	case "PUT":
		responces.Put(db, dbname, newcnt, idInt)
	case "DELETE":
		responces.Delete(db, dbname, idInt)
	case "POST":
		responces.Post(db, dbname, newcnt)
	default:
		log.Fatalln("Undefined method")
	}
}
