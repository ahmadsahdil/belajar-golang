package golang_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(mysql-server:3306)/golang")
	if err != nil {
		panic(err)

	} else {
		fmt.Println("berhasil")
	}

	defer db.Close()

}
