// Package router faz as configurações das rotas para a API
package router

import (
	"weni/api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna um router com as rotas configuradas.
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
