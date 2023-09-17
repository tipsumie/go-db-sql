package repository

import (
	"database/sql"
	"go-db-sql/models"
)

type CustomerRepo struct {
	db *sql.DB
}

func ConnectCustomerRepository(database *sql.DB) *CustomerRepo {
	return &CustomerRepo{db: database}
}

func (cr *CustomerRepo) GetCustomers() ([]models.Customer, error) {
	err := cr.db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select * from customers"
	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}

	// create customers slice
	customers := []models.Customer{}

	// read the value of each row.
	for rows.Next() {
		customer := models.Customer{}
		err = rows.Scan(&customer.CustomerId, &customer.FirstName, &customer.LastName,
			&customer.Email, &customer.PhoneNumber, &customer.Address, &customer.City,
			&customer.PostalCode, &customer.Country)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
