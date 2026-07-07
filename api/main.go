package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println("Iniciando API...")
	// fmt.Println(config.Porta, config.StringConexaoBanco)
	fmt.Printf("Escutando na porta %d \n", config.Porta)
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
