package responces

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/voyadger01/rest-contacts_golang_01/contact"
)

func isThereId(db *sql.DB, dbname string, cnt contact.Contact, id int) bool {
	rows, err := db.Query("SELECT * FROM ", dbname)
	if err != nil {
		log.Fatalln("bad num of rows")
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Fatalln("bad num of columns")
	}
	isId := false
	for i := 0; i < len(cols); i++ {
		colsName, err := strconv.Atoi(cols[i])
		if err != nil {
			log.Fatalln("bad format of cols")
		}
		if cnt.Id == colsName {
			isId = true
		}
	}
	return isId
}

func Put(db *sql.DB, dbname string, cnt contact.Contact, id int) {
	isId := isThereId(db, dbname, cnt, id)
	if isId {
		Delete(db, dbname, id)
		Post(db, dbname, cnt)
	} else {
		Post(db, dbname, cnt)
	}
}
