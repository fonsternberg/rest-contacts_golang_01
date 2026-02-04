package responces

import (
	"database/sql"
	"log"
	"github.com/voyadger01/rest-contacts_golang_01/contact"
)

func GetOne(db *sql.DB, dbname string, id int) {
	rows, err := db.Query("SELECT ", id, "from ", dbname)
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
}