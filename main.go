package main

import (
	"database/sql"
	"fmt"
	"go-db-sql/config"
	"go-db-sql/repository"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

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

	// mysql dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Database)

	// postgres dsn
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.Database.Host,
	// 	cfg.Database.Username,
	// 	cfg.Database.Password,
	// 	cfg.Database.Database)

	db, err := sql.Open(cfg.Database.Driver, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := repository.ConnectCustomerRepository(db)

	// customers, err := repo.GetCustomers()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%#v", customers)

	customer, err := repo.GetCustomer(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", customer)

}
