package golang_database

import (
	"testing"
	"context"
	"fmt"
)

func TestExecSql(t *testing.T)  {
	db := GetConnection()
	defer db.Close()


	ctx := context.Background()
	query := "INSERT INTO customer(id, name) VALUES('muhammad','muhammad')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert into database")
	
}

func TestQuerySql(t *testing.T)  {
	db := GetConnection()
	defer db.Close()


	ctx := context.Background()
	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID: ",id )
		fmt.Println("Name: ",name )
	}
	
}