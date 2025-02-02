package service

import (
	"banking/domain"
	"banking/dto"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, error)
	GetAllCustomerById(id string) (dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, error) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}
func (s DefaultCustomerService) GetAllCustomerById(id string) (dto.CustomerResponse, error) {
	c, _ := s.repo.FindById(id)
	//if err != nil {
	//	return nil, err
	//}
	response := s.repo.ToDTO(c)
	return response, nil
}
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repository,
	}
}
