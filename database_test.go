package learn_golang_mysql_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDatabase(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/learn-golang-mysql-database")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
