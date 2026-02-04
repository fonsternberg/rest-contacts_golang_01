package responces

import (
	"database/sql"
	"log"
)

func GetOne(db *sql.DB, dbname string, id int) {
	rows, err := db.Query("SELECT ", id, "from ", dbname)
	if err != nil {
		log.Fatalln("bad query")		
	}
	
}