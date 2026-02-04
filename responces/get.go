package responces

import (
	"database/sql"
	"fmt"
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
	for i := range contacts {
		fmt.Println("Contact ", id, ": ", contacts[i].Name, " ", contacts[i].Phone)
	}
}

func GetAll(db *sql.DB, dbname string) {
	count := 0
	err := db.QueryRow("SELECT COUNT(*) FROM ", dbname).Scan(&count)
	if err != nil {
		log.Println("Zero elements")
		return
	}
	for i := 1; i < count; i++ {
		GetOne(db, dbname, i)
	}
}