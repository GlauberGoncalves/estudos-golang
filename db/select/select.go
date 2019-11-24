package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	ID   int
	Nome string
}

func main() {

	db, err := sql.Open("mysql", "root:12345678@/cursogo")
	if err != nil {
		fmt.Errorf("erro ao conectar ao banco de dados")
	}

	defer db.Close()

	rows, err := db.Query("select id, nome from usuarios where id < ?", 5)
	if err != err {
		fmt.Errorf("erro ao retornar query")
	}
	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome)
		usuarios = append(usuarios, u)
	}

	fmt.Println(usuarios)

}
