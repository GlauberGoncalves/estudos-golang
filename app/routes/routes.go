package routes

import (
	"estudos-golang/app/controllers"
	"net/http"
)

// LoadRoutes Carrega todas as notas do sistema
func LoadRoutes() {

	http.HandleFunc("/minhas-notas", controllers.MinhasNotas)
	http.HandleFunc("/nova", controllers.NovaNota)
}
