package modelos

// DadosAutenticacao contem o token e o id do usuario logado/autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
