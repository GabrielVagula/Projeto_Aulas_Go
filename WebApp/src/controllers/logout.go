package controllers

import (
	"WebApp/src/cookies"
	"net/http"
)

// FazerLogout remove os dados de autenticação salvos no browser do usuario.
func FazerLogout(w http.ResponseWriter, r *http.Request) {

	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}
