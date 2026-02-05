package responces

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/voyadger01/rest-contacts_golang_01/contact"
)

func isThereId(db *sql.DB, dbname string, cnt contact.Contact) bool {
	rows, err := db.Query("SELECT * FROM ", dbname)
	if err != nil {
		log.Fatalln("bad num of rows")
	}
	isId := false
	for rows.Next() {
		var tempID string
		err := rows.Scan(tempID)
		if err != nil {
			log.Fatalln("Wrong Id format")
		}
		currentId, err := strconv.Atoi(tempID)
		if err != nil {
			log.Fatalln("id must be integer")
		}
		if cnt.Id == currentId {
			isId = true
			break
		}
	}
	return isId
}

func Put(db *sql.DB, dbname string, cnt contact.Contact, id int) {
	isId := isThereId(db, dbname, cnt)
	if isId {
		Delete(db, dbname, id)
		Post(db, dbname, cnt)
	} else {
		Post(db, dbname, cnt)
	}
}
