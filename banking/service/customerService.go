package service

import (
	"github.com/vairarchi/rest-microservices-api-course/banking/domain"
	"github.com/vairarchi/rest-microservices-api-course/banking/dto"
	"github.com/vairarchi/rest-microservices-api-course/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	// fmt.Println("status from query param:", status)
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	var response []dto.CustomerResponse
	for _, v := range c {
		response = append(response, v.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
