package models

type Customer struct {
	CustomerId  int    `db:"customer_id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	Address     string `db:"address"`
	City        string `db:"city"`
	PostalCode  string `db:"postal_code"`
	Country     string `db:"country"`
}
