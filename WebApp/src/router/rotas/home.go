package rotas

import (
	"WebApp/src/controllers"
	"net/http"
)

var rotaPaginaPrincipal = Rota{
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
