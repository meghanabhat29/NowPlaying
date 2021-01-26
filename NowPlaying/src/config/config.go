package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetDB ...
func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "-root"
	dbPass := "Intern@paytm29"
	dbName := "Playlist"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
