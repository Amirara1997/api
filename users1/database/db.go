package database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() (*gorm.DB, error) {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 "admin",
		Passwd:               "Am!r13769900Am!r",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "dorf",
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	return gorm.Open(DBMS, mySqlConfig.FormatDSN())
}