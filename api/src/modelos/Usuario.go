package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	Id       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Valida e formata o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	usuario.formatar()

	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("Nome é obrigatório.")
	}

	if usuario.Nick == "" {
		return errors.New("Nick é obrigatório.")
	}

	if usuario.Email == "" {
		return errors.New("E-mail é obrigatório.")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("Senha é obrigatória.")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
