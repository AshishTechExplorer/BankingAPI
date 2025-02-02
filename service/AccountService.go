package service

import (
	"banking/domain"
	"banking/dto"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, error)
}
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, error) {
	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	newAccount, _ := s.repo.Save(account)
	return newAccount.ToNewAccountResponseDto(), nil
}
