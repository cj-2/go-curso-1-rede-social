package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar irá retornar um router com as rotas configuradas.
func Gerar() *mux.Router {
	r := mux.NewRouter()
	rotas.Configurar(r)
	return r
}
