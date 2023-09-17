package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/customer")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	query := "select customer_id, first_name, last_name from customers"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	// read value each rows.
	for rows.Next() {
		customer_id := 0
		first_name := ""
		last_name := ""

		err = rows.Scan(&customer_id, &first_name, &last_name)
		if err != nil {
			panic(err)
		}

		fmt.Println(customer_id, first_name, last_name)
	}
}
