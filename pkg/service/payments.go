package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
)

type PaymentsService struct {
	repo repository.Payments
}

func NewPaymentsService(repo repository.Payments) *PaymentsService {
	return &PaymentsService{repo: repo}
}

func (s *PaymentsService) Create(payment taskexchange.Payment) (int, error) {
	return s.repo.Create(payment)
}

func (s *PaymentsService) GetByUser(user taskexchange.User) ([]taskexchange.Payment, error) {
	return s.repo.GetByUser(user)
}
