package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	Result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return Result
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	exec(db, "Create database if not exist cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exist usuarios")
	exec(db, `create table usuarios (
    id integer auto_increment,
    nome varchar(80),
    PRIMARY KEY (id),
    )`)
}
