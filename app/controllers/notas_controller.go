package controllers

import (
	"Go-JWT-Postgres-Mysql-Restful-API/api/responses"
	"Go-JWT-Postgres-Mysql-Restful-API/api/utils/formaterror"
	"estudos-golang/app/config"
	"estudos-golang/app/models"
	"fmt"
	"log"
	"net/http"
)

func NovaNota(w http.ResponseWriter, r *http.Request) {

	db, _ := config.Connect()

	var nota models.Nota

	fmt.Println(r.FormValue("titulo"))

	nota.Titulo = r.FormValue("titulo")
	nota.Conteudo = r.FormValue("conteudo")

	notaCreated, err := nota.SalvaNota(db)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, notaCreated.ID))
	defer db.Close()
	responses.JSON(w, http.StatusCreated, notaCreated)
}

func MinhasNotas(w http.ResponseWriter, r *http.Request) {

	nota := models.Nota{}
	db, err := config.Connect()
	if err != nil {
		log.Println("Não foi possivel conectar ao banco de dados")
	}

	notas, err := nota.AllNotas(db)

	w.Header().Set("Content-Type", "application/json")
	responses.JSON(w, http.StatusOK, notas)
}
