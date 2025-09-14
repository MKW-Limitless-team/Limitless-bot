package globals

import (
	"database/sql"
	"log"
	"os"
)

var (
	TOKEN string
	DB    *sql.DB
)

func Init() {
	TOKEN = os.Getenv("BEPIS_TOKEN")
	DB, err := sql.Open("sqlite3", "./ltrc.db")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
