package golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	
)


func TestOpenConnection(t *testing.T)  {
	db, err := sql.Open("mysql", "root:password@tcp(mysql-server:3306)/golang")
	if err != nil {
		panic(err)
	} 


	defer db.Close()

	
}