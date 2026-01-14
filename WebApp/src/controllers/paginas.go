package controllers

import (
	"WebApp/src/config"
	"WebApp/src/requisicoes"
	"WebApp/src/utils"
	"fmt"
	"net/http"
)

// CarregarTelaDeLogin vai renderezar a tela de Login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario vai carregar a página de cadastro de usuario
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal carrega a pagina principal com as publicações.
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)

	// 1. Verifique o erro PRIMEIRO
	if erro != nil {
		fmt.Println("Erro ao fazer a requisição:", erro)
		utils.ExecutarTemplate(w, "home.html", nil) // Ou uma página de erro
		return
	}

	// 2. Garanta que o corpo da resposta será fechado
	defer response.Body.Close()

	// 3. Agora sim você pode acessar o StatusCode com segurança
	fmt.Println("Status Code da API:", response.StatusCode, "Erro:", erro)

	// Aqui você provavelmente vai querer tratar se o status não for 200
	if response.StatusCode >= 400 {
		fmt.Println("A API retornou um erro:", response.StatusCode)
	}

	utils.ExecutarTemplate(w, "home.html", nil)
}
