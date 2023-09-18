package main

import (
	"fmt"
	"go-db-sql/config"
	"go-db-sql/models"
	"go-db-sql/repository"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

	db, err := sqlx.Open(cfg.Database.Driver, dsn)
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

	// get data by id
	// customer, err := repo.GetCustomer(1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%#v", customer)

	// insert data
	// newCustomer := models.Customer{
	// 	CustomerId:  21,
	// 	FirstName:   "Mimie",
	// 	LastName:    "Poko",
	// 	Email:       "test@gmail.com",
	// 	PhoneNumber: "099999999",
	// 	Address:     "The Earth",
	// 	City:        "BK",
	// 	PostalCode:  "5555",
	// }
	// err = repo.AddCustomer(newCustomer)
	// if err != nil {
	// 	log.Fatalf("Failed to insert data: %v", err)
	// }

	// update data
	updateCustomer := models.Customer{
		CustomerId:  21,
		FirstName:   "Mimi",
		LastName:    "Poko",
		Email:       "test@gmail.com",
		PhoneNumber: "099999999",
		Address:     "The Earth",
		City:        "BK",
		PostalCode:  "5555",
	}
	err = repo.UpdateCustomer(updateCustomer)
	if err != nil {
		log.Fatalf("Failed to update data: %v", err)
	}

	// delete data
	// err = repo.DeleteCustomer(21)
	// if err != nil {
	// 	log.Fatalf("Failed to delete data: %v", err)
	// }
}
