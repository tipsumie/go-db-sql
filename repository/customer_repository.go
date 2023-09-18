package repository

import (
	"database/sql"
	"errors"
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

func (cr *CustomerRepo) GetCustomer(id int) (*models.Customer, error) {
	err := cr.db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select * from customers where customer_id=?"
	row := cr.db.QueryRow(query, id)
	customer := models.Customer{}
	err = row.Scan(&customer.CustomerId, &customer.FirstName, &customer.LastName, &customer.Email,
		&customer.PhoneNumber, &customer.Address, &customer.City, &customer.Country, &customer.PostalCode)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (cr *CustomerRepo) AddCustomer(customer models.Customer) error {
	err := cr.db.Ping()
	if err != nil {
		return err
	}

	query := "insert into customers (customer_id, first_name, last_name, email, phone_number, address, city, country, postal_code) values(?,?,?,?,?,?,?,?,?)"
	result, err := cr.db.Exec(query, customer.CustomerId, customer.FirstName, customer.LastName, customer.Email, customer.PhoneNumber, customer.Address, customer.City, customer.Country, customer.PostalCode)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Can't insert data")
	}
	return nil
}

func (cr *CustomerRepo) UpdateCustomer(customer models.Customer) error {
	err := cr.db.Ping()
	if err != nil {
		return err
	}

	query := "update customers set first_name=? where customer_id=?"
	result, err := cr.db.Exec(query, customer.FirstName, customer.CustomerId)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Can't update data")
	}
	return nil
}

func (cr *CustomerRepo) DeleteCustomer(id int) error {
	err := cr.db.Ping()
	if err != nil {
		return err
	}

	query := "delete from customers where customer_id=?"
	result, err := cr.db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return errors.New("Can't delete data")
	}
	return nil
}
