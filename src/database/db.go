package database

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

type DBInstance struct {
	DB *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS url (
    id SERIAL,
    original TEXT NOT NULL UNIQUE,
    key VARCHAR(10) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);
`

func (di *DBInstance) Connect() {
	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Fatalln("Database url (DB_URL) not specified.")
	}

	for range time.Tick(5 * time.Second) {
		log.Println("Trying to connect to database...")

		db, err := sqlx.Connect("pgx", dbUrl)
		if err == nil {
			di.DB = db

			_, err = di.DB.Exec(schema)
			if err != nil {
				log.Fatalln(err)
			}

			return
		}

		log.Println(err)
	}
}
