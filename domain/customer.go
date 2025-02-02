package domain

import "banking/dto"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, error)
	FindById(id string) (*Customer, error)
	ToDTO(c *Customer) dto.CustomerResponse // * is used to return nil in case no customer,nil is correct
}

func (d CustomerRepositoryDb) ToDTO(c *Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.Status,
	}
}
