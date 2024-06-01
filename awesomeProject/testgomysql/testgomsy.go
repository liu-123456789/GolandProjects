package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 使用正确的用户名、密码和IP地址
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/webook")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	println("Connected to the database successfully!")
}
