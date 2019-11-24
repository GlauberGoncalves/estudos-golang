package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite3"
)

type Usuario struct {
	ID    int    `gorm:"primary_key;auto_increment" json:"id"`
	Nome  string `gorm:"size:255;not null;unique" json:"nome"`
	Senha string `gorm:"size:255;not null;unique" json:"senha"`
}

func main() {

	db, err := gorm.Open("sqlite3", ":memory:")
	defer db.Close()

	if err != nil {
		fmt.Errorf("Não foi possivel criar banco de dados")
	}

	usuario1 := Usuario{ID: 1, Nome: "Glauber", Senha: "123456"}
	re := db.Save(&usuario1)

	fmt.Print(re)

	var usuario Usuario
	fmt.Println(usuario)

	db.First(&usuario, 1)

	fmt.Println(usuario)

}
