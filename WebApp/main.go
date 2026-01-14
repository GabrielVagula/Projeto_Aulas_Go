package main

import (
	"WebApp/src/config"
	"WebApp/src/cookies"
	"WebApp/src/router"
	"WebApp/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
