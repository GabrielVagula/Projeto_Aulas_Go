package controllers

import (
	"API/src/autenticacao"
	"API/src/banco"
	"API/src/modelos"
	"API/src/repositorios"
	"API/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarPublicacao adiciona uma nova publicação no banco de dados.
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)

}

// BuscarPublicacoes traz as publicações que aparecem no feed do usuario.
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

// BuscarPublicacao busca uma publicação especifica de algum outro usuario.
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// AtualizarPublicacao atualiza algo na publicação que o usuario deseja alterar.
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// DeletarPublicacao deleta do banco de dados uma publicação que o usuario quer.
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
