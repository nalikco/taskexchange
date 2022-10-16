package service

import (
	"database/sql"
	"errors"
	"fmt"
	"taskexchange"
	"taskexchange/pkg/repository"
	"time"
)

type OffersService struct {
	offersRepo repository.Offers
	tasksRepo  repository.Tasks
	usersRepo  repository.Users
	eventsRepo repository.Events
	ordersRepo repository.Orders
}

func NewOffersService(offersRepo repository.Offers, tasksRepo repository.Tasks, usersRepo repository.Users, eventsRepo repository.Events, ordersRepo repository.Orders) *OffersService {
	return &OffersService{
		offersRepo: offersRepo,
		tasksRepo:  tasksRepo,
		usersRepo:  usersRepo,
		eventsRepo: eventsRepo,
		ordersRepo: ordersRepo,
	}
}

func (s *OffersService) GetPerformerActive(performerId int) ([]taskexchange.Offer, error) {
	return s.offersRepo.FindPerformerActiveOffers(performerId)
}

func (s *OffersService) Make(performerId, taskId int) (int, error) {
	task, err := s.tasksRepo.GetById(taskId)
	if err != nil {
		return 0, err
	}
	if task.Status != 1 {
		return 0, errors.New("wrong task status")
	}

	taskOrders, err := s.ordersRepo.FindActiveByTaskId(taskId)

	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if err == nil {
		for _, order := range taskOrders {
			orderOffer, err := s.offersRepo.FindOneById(order.OfferId)
			if err != nil {
				return 0, err
			}

			if orderOffer.PerformerId == performerId {
				return 0, errors.New("active order for this task already exists")
			}
		}
	}

	id, err := s.offersRepo.Create(performerId, taskId)
	if err != nil {
		return 0, err
	}

	_, _ = s.eventsRepo.Create(taskexchange.Event{
		UserId:    task.CustomerId,
		Message:   fmt.Sprintf("Новое предложение по задаче #%d от исполнителя ID:%d", task.Id, performerId),
		Link:      "/tasks/my",
		CreatedAt: time.Now(),
	})

	return id, nil
}

func (s *OffersService) ChangeStatus(offerId, customerId, status int) error {
	offer, err := s.offersRepo.FindOneById(offerId)
	if err != nil {
		return err
	}
	if offer.Status != 0 {
		return errors.New("wrong offer status")
	}
	task, err := s.tasksRepo.GetById(offer.TaskId)
	if err != nil {
		return err
	}
	if task.CustomerId != customerId {
		return errors.New("access denied")
	}

	err = s.offersRepo.Update(offerId, taskexchange.UpdateOfferInput{
		Status: status,
	})
	if err != nil {
		return err
	}

	if status == 1 {
		newTaskAmount := task.Amount - 1
		updateTaskInput := taskexchange.UpdateTaskInput{
			Amount: &newTaskAmount,
		}
		if newTaskAmount < 1 {
			newStatus := 0
			updateTaskInput.Status = &newStatus
		}
		err = s.tasksRepo.Update(task.Id, updateTaskInput)
		orderId, err := s.ordersRepo.Create(offerId, task.Id)
		if err != nil {
			return err
		}

		_, err = s.eventsRepo.Create(taskexchange.Event{
			UserId:    offer.PerformerId,
			Message:   fmt.Sprintf("Заказчик принял Ваше предложение по задаче #%d. Заказ #%d создан.", task.Id, orderId),
			Link:      "/orders/performer",
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}
	} else {
		_, err = s.eventsRepo.Create(taskexchange.Event{
			UserId:    offer.PerformerId,
			Message:   fmt.Sprintf("Заказчик отклонил Ваше предложение по задаче #%d", task.Id),
			Link:      "/orders/performer",
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
