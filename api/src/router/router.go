package router

import (
	"API/src/router/rotas/rotas"

	"github.com/gorilla/mux"
)

// Essa func gerar vai retornar um router com as rotas configuradas.
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)

}
