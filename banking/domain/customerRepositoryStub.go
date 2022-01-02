package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1002", Name: "Santri", City: "Nagpur", Zipcode: "440023", DateofBirth: "2002-01-01", Status: "1"},
		{Id: "1001", Name: "Vaibhav", City: "Pune", Zipcode: "411014", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1003", Name: "Pratik", City: "Nagpur", Zipcode: "440024", DateofBirth: "2003-01-01", Status: "1"},
		{Id: "1004", Name: "Mandal", City: "Pune", Zipcode: "411014", DateofBirth: "2004-01-01", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
