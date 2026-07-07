package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	// Copiando o conteúdo de body para usuário.
	var usuario modelos.Usuario
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		log.Fatal(erro)
	}

	// Conexão com banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	// Criação do repositório:
	repositorio := repositorios.NovoRespositorioDeUsuarios(db)

	id, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Criando Usuário %d", id)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuários"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário"))
}

func ApagarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Apagando Usuário"))
}
