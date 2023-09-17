package main

import (
	"database/sql"
	"fmt"
	"go-db-sql/config"
	"go-db-sql/repository"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// var err error
	// db, err = sql.Open("mysql", "root:password@tcp(localhost)/customer")
	// if err != nil {
	// 	panic(err)
	// }

	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Database)

	fmt.Printf("Database Config: %+v\n", cfg.Database)
	fmt.Println("DSN:", dsn)

	// dsn := "root:password@tcp(localhost)/customer"

	db, err := sql.Open(cfg.Database.Driver, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := repository.ConnectCustomerRepository(db)
	customers, err := repo.GetCustomers()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v", customers)

}
