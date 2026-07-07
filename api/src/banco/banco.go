package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Abre conexão com o banco de dados e retorna.
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close() // Fechamos a conexão garantir que nada se mantenha aberto.
		return nil, erro
	}

	return db, nil
}
