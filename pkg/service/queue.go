package service

import (
	"taskexchange/pkg/repository"
	"time"
)

type QueueService struct {
	tasks  repository.Tasks
	orders repository.Orders
}

func NewQueueService(tasks repository.Tasks, orders repository.Orders) *QueueService {
	return &QueueService{
		tasks:  tasks,
		orders: orders,
	}
}

func (s *QueueService) Run() error {
	for true {

		time.Sleep(time.Minute)
	}
	return nil
}
