package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func Dbinit() {
	var err error
	Db, err = sql.Open("mysql", "root:jly0614@tcp(127.0.0.1:3306)/homework")
	//Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/homework")

	if err != nil {
		log.Panicln("err:", err.Error())
	}
	fmt.Print("连接成功")
	Db.SetMaxOpenConns(10)   //设置与数据库建立连接的最大数目
	Db.SetMaxIdleConns(10)	//设置连接池中的最大闲置连接数
}
