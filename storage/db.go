package storage

import (
	"example.com/gin_realworld/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sqlx.DB
var gormDb *gorm.DB

func init() {
	var err error
	//dsn := "root:12345678@tcp(127.0.0.1:3307)/realworld?parseTime=true"
	db, err = sqlx.Open("mysql", config.GetMySQLDSN())
	//db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	gormDb, err = gorm.Open(gorm_mysql.New(gorm_mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	err = gormDb.Exec("select 1;").Error
	if err != nil {
		panic(err)
	}
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
