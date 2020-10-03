package conf

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

func NewDatabaseConnection() (*sql.DB, error) {
	source := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", "db", "postgres", "devdb", "postgres")
	conn, err := sql.Open("postgres", source)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
