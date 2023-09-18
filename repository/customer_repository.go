package repository

import (
	"go-db-sql/models"

	"github.com/jmoiron/sqlx"
)

type CustomerRepo struct {
	db *sqlx.DB
}

func ConnectCustomerRepository(database *sqlx.DB) *CustomerRepo {
	return &CustomerRepo{db: database}
}

func (cr *CustomerRepo) GetCustomers() ([]models.Customer, error) {
	query := "select * from customers"
	customers := []models.Customer{}
	err := cr.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (cr *CustomerRepo) GetCustomer(id int) (*models.Customer, error) {
	query := "select * from customers where customer_id=:id"
	customer := models.Customer{}
	err := cr.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (cr *CustomerRepo) AddCustomer(customer models.Customer) error {
	tx, err := cr.db.Beginx()
	if err != nil {
		return err
	}

	query := `
		insert into customers (customer_id, first_name, last_name, email, phone_number, address, city, country, postal_code) 
		values(:customer_id, :first_name, :last_name, :email, :phone_number, :address, :city, :country, :postal_code)`
	_, err = tx.NamedExec(query, customer)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (cr *CustomerRepo) UpdateCustomer(customer models.Customer) error {
	tx, err := cr.db.Beginx()
	if err != nil {
		return err
	}

	query := `update customers set first_name=:first_name where customer_id=:customer_id`
	_, err = tx.NamedExec(query, customer)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (cr *CustomerRepo) DeleteCustomer(id int) error {
	tx, err := cr.db.Beginx()
	if err != nil {
		return err
	}

	query := "delete from customers where customer_id=:id"
	_, err = tx.NamedExec(query, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
