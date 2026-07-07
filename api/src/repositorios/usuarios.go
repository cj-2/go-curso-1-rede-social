package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// Cria um repositório de usuários
func NovoRespositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar usuário no banco de dados.
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
