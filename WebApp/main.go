package main

import (
	"WebApp/src/router"
	"WebApp/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {

	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Rodando WebApp")
	log.Fatal(http.ListenAndServe(":3000", r))

}
