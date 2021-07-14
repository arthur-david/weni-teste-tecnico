// Package main vai chamar o package config e o package router
package main

import (
	"fmt"
	"log"
	"net/http"
	"weni/api/src/config"
	"weni/api/src/router"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
