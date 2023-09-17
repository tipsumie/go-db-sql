package main

import (
	"database/sql"
	"fmt"
	"go-db-sql/repository"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost)/customer")
	if err != nil {
		panic(err)
	}

	repo := repository.ConnectCustomerRepository(db)
	customers, err := repo.GetCustomers()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", customers)

}
