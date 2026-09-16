package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa uma rota da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	routas := rotasUsuarios
	routas = append(routas, rotaLogin)

	for _, rota := range routas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
