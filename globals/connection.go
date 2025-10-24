package globals

import "database/sql"

var (
	DB         *sql.DB
	SQLITEFILE = "./ltrc.db"
)

func Initialize(filePath string) error {
	db, err := sql.Open("sqlite3", SQLITEFILE)
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func GetConnection() *sql.DB {
	err := DB.Ping()
	if err != nil {
		return DB
	}

	Initialize(SQLITEFILE)
	return DB
}
