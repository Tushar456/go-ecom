package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase() (*Database, error) {

	db, err := sqlx.Open("mysql", "root:bunty@123@tcp(localhost:3306)/ecom?parseTime=true")

	if err != nil {

		return nil, err
	}

	return &Database{db: db}, nil

}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetDb() *sqlx.DB {
	return d.db
}
