package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Conectar Abre a conexão com o banco de dados
func Conectar()(*sql.DB, error)  {
	stringConexao := "root:Password123-@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, error := sql.Open("mysql", stringConexao)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
