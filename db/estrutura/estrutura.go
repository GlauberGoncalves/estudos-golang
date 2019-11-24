package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {

	db, err := sql.Open("mysql", "root:12345678@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `
		create table usuarios(
			id integer not null auto_increment primary key,
			nome varchar(255)			
		)
	`)

	exec(db, "insert into usuarios(id, nome) values (default, 'glauber')")

}
