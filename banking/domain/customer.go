package domain

import "github.com/vairarchi/rest-microservices-api-course/banking/errs"

type Customer struct {
	Id          string `json:"customer_id" db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	// status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
