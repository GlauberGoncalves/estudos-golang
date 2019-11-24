package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Usuario struct {
	ID    int    `gorm:"default:'id'"`
	Nome  string `gorm:"default:'nome'"`
	Senha string `gorm:"default:'senha'"`
}

func main() {

	db, err := gorm.Open("sqlite3", "/test.db")
	defer db.Close()

	if err != nil {
		fmt.Errorf("Não foi possivel criar banco de dados")
	}

	usuario := Usuario{ID: 1, Nome: "Glauber", Senha: "123456"}
	db.NewRecord(usuario) // => returns `true` as primary key is blank

}
