package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewMysqlHandler() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "linqumate"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	fmt.Println(CONNECT)

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println("エラー")
		panic(err.Error())
	}
	return db
}
