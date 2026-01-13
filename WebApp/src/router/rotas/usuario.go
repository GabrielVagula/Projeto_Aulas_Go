package rotas

import (
	"WebApp/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuarios,
		RequerAutenticacao: false,
	},
}
