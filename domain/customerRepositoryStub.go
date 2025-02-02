package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Ashish", City: "New Delhi", DateOfBirth: "2000-10-10",
			Status: "1", Zipcode: "201301"},
		{Id: "1002", Name: "Robin", City: "New Delhi", DateOfBirth: "2000-10-10",
			Status: "1", Zipcode: "201301"},
	}
	return CustomerRepositoryStub{customers: customers}
}
