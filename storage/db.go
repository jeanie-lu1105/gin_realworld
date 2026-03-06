package storage

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init(database *sqlx.DB){
	db = sqlx.Open("mysql", ":memory:")
}