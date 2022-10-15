package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
)

type OrdersService struct {
	ordersRepo repository.Orders
}

func NewOrdersService(ordersRepo repository.Orders) *OrdersService {
	return &OrdersService{
		ordersRepo: ordersRepo,
	}
}

func (s *OrdersService) FindActiveByPerformerId(performerId int) ([]taskexchange.Order, error) {
	return s.ordersRepo.FindActiveByPerformerId(performerId)
}
