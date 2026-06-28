package router

import "github.com/gorilla/mux"

// Gerar irá retornar um router com as rotas configuradas.
func Gerar() *mux.Router {
	return mux.NewRouter()
}
