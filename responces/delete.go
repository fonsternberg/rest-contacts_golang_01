package responces

import (
	"database/sql"
	"log"
)

func delete(db *sql.DB, dbname string, id int) {
	res, err := db.Exec("DELETE FROM ", dbname, "WHERE ", "id = ", id)
	if err != nil {
		log.Fatalln("bad id to delete")
	}
	rowsDeleted, _ := res.RowsAffected()
	log.Println("Rows deleted: ", rowsDeleted)
}
