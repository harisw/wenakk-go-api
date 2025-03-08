package helpers

import (
	"database/sql"
	"log"
)

func HandleSQLError(err error, msg string) bool {
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found")
		} else {
			log.Println(msg+": Query error:", err)
		}
		return true
	}
	return false
}
