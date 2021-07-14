// Package database é o pacote reponsável por criar e testar uma nova conexão com o banco de dados.
package database

import (
	"database/sql"
	"weni/api/src/config"

	_ "github.com/go-sql-driver/mysql" //Driver
)

// Conectar abre a conexão com o banco de dados e retorna a conexão.
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro := db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil

}
