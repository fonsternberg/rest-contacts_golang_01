package responces

import (
	"database/sql"
	"errors"
)

func Execute(db *sql.DB, responce string) (sql.Result, error) {
	res, err := db.Exec(responce)
	if err != nil {
		return nil, errors.New("can't execute")
	}
	return res, err
}