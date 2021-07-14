// Package responses retorna uma resposta JSON para uma requisição. Retorna um JSON de resposta ou um JSON de erro.
package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON para a requisição.
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json") // Define o tipo de conteudo da resposta como JSON.
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Erro retorna um erro em formato JSON para a requisição, caso haja um erro.
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(), // Preenchendo struct que foi passado como parametro da chamada de função JSON
	})
}
