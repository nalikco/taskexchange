package service

import (
	"errors"
	"fmt"
	"taskexchange"
	"taskexchange/pkg/repository"
	"time"
)

type OrdersService struct {
	ordersRepo      repository.Orders
	usersRepo       repository.Users
	optionsRepo     repository.Options
	eventsRepo      repository.Events
	tasksRepo       repository.Tasks
	taskOptionsRepo repository.TaskOptions
	paymentsRepo repository.Payments
}

func NewOrdersService(ordersRepo repository.Orders, usersRepo repository.Users, optionsRepo repository.Options, tasksRepo repository.Tasks, taskOptionsRepo repository.TaskOptions, eventsRepo repository.Events, paymentsRepo repository.Payments) *OrdersService {
	return &OrdersService{
		ordersRepo:      ordersRepo,
		usersRepo:       usersRepo,
		optionsRepo:     optionsRepo,
		tasksRepo:       tasksRepo,
		eventsRepo:      eventsRepo,
		taskOptionsRepo: taskOptionsRepo,
		paymentsRepo: paymentsRepo,
	}
}

func (s *OrdersService) FindAllByPerformerId(performerId int) ([]taskexchange.Order, error) {
	orders, err := s.ordersRepo.FindAllByPerformerId(performerId)
	if err != nil {
		return []taskexchange.Order{}, err
	}

	for i, order := range orders {
		orders[i].Task.Customer, err = s.usersRepo.GetById(order.Task.CustomerId, false)
		if err != nil {
			return []taskexchange.Order{}, err
		}

		taskOptions, err := s.taskOptionsRepo.GetByTaskId(order.Task.Id)
		if err != nil {
			return []taskexchange.Order{}, err
		}
		var taskOptionsIds []int

		for _, taskOption := range taskOptions {
			taskOptionsIds = append(taskOptionsIds, taskOption.OptionId)
		}

		options, err := s.optionsRepo.GetByIds(taskOptionsIds)
		if err != nil {
			return []taskexchange.Order{}, err
		}

		orders[i].Task.Options = options
	}

	return orders, nil
}

func (s *OrdersService) FindAllByCustomerId(customerId int) ([]taskexchange.Order, error) {
	orders, err := s.ordersRepo.FindAllByCustomerId(customerId)
	if err != nil {
		return []taskexchange.Order{}, err
	}

	for i, order := range orders {
		orders[i].Offer.Performer, err = s.usersRepo.GetById(order.Offer.PerformerId, false)
		if err != nil {
			return []taskexchange.Order{}, err
		}

		taskOptions, err := s.taskOptionsRepo.GetByTaskId(order.Task.Id)
		if err != nil {
			return []taskexchange.Order{}, err
		}
		var taskOptionsIds []int

		for _, taskOption := range taskOptions {
			taskOptionsIds = append(taskOptionsIds, taskOption.OptionId)
		}

		options, err := s.optionsRepo.GetByIds(taskOptionsIds)
		if err != nil {
			return []taskexchange.Order{}, err
		}

		orders[i].Task.Options = options
	}

	return orders, nil
}

func (s *OrdersService) FindAll() ([]taskexchange.Order, error) {
	orders, err := s.ordersRepo.FindAll()
	if err != nil {
		return []taskexchange.Order{}, err
	}

	for i, order := range orders {
		orders[i].Task.Customer, err = s.usersRepo.GetById(order.Task.CustomerId, false)
		if err != nil {
			return []taskexchange.Order{}, err
		}

		orders[i].Offer.Performer, err = s.usersRepo.GetById(order.Offer.PerformerId, false)
		if err != nil {
			return []taskexchange.Order{}, err
		}

		taskOptions, err := s.taskOptionsRepo.GetByTaskId(order.Task.Id)
		if err != nil {
			return []taskexchange.Order{}, err
		}
		var taskOptionsIds []int

		for _, taskOption := range taskOptions {
			taskOptionsIds = append(taskOptionsIds, taskOption.OptionId)
		}

		options, err := s.optionsRepo.GetByIds(taskOptionsIds)
		if err != nil {
			return []taskexchange.Order{}, err
		}

		orders[i].Task.Options = options
	}

	return orders, nil
}

func (s *OrdersService) FindActiveByPerformerId(performerId int) ([]taskexchange.Order, error) {
	return s.ordersRepo.FindActiveByPerformerId(performerId)
}

func (s *OrdersService) Update(orderId int, userId int, input taskexchange.UpdateOrderInput) error {
	user, err := s.usersRepo.GetById(userId, true)
	if err != nil {
		return err
	}

	if user.Type != 3 && input.TaskId != nil {
		*input.TaskId = 0
	}

	order, err := s.ordersRepo.FindOneById(orderId)
	if err != nil {
		return err
	}

	if order.Task.CustomerId != user.Id && order.Offer.PerformerId != user.Id && user.Type != 3 {
		return errors.New("access denied")
	}

	if input.Status != nil {
		if order.Status == 1 {
			if order.Task.CustomerId == user.Id {
				if *input.Status == 0 {
					if input.ReturnComment == nil {
						return errors.New("wrong return comment")
					}

					_, _ = s.eventsRepo.Create(taskexchange.Event{
						UserId:    order.Offer.PerformerId,
						Message:   fmt.Sprintf("Заказчик вернул заказ #%d в работу.", order.Id),
						Link:      "/orders/performer",
						CreatedAt: time.Now(),
					})

					return s.ordersRepo.Update(orderId, taskexchange.UpdateOrderInput{
						Status:        input.Status,
						ReturnComment: input.ReturnComment,
					})
				}
				if *input.Status == 2 {
					taskOptions, err := s.taskOptionsRepo.GetByTaskId(order.Task.Id)
					if err != nil {
						return err
					}
					var taskOptionsIds []int

					for _, taskOption := range taskOptions {
						taskOptionsIds = append(taskOptionsIds, taskOption.OptionId)
					}

					options, err := s.optionsRepo.GetByIds(taskOptionsIds)
					if err != nil {
						return err
					}

					order.Task.Options = options

					performer, err := s.usersRepo.GetById(order.Offer.PerformerId, true)
					if err != nil {
						return err
					}
					newPerformerBalance := performer.Balance + order.Task.CalculatePriceForOne()

					err = s.usersRepo.Update(performer.Id, taskexchange.UpdateUserInput{
						Balance: &newPerformerBalance,
					})
					if err != nil {
						return err
					}

					payment := taskexchange.Payment{
						User: performer,
						Type: 2,
						Comment: "Выполнение задач",
						Sum: order.Task.CalculatePriceForOne(),
					}
					_, err = s.paymentsRepo.Create(payment)
					if err != nil {
						return err
					}

					_, _ = s.eventsRepo.Create(taskexchange.Event{
						UserId:    order.Offer.PerformerId,
						Message:   fmt.Sprintf("Заказчик подтвердил выпонение заказа #%d.", order.Id),
						Link:      "/orders/performer",
						CreatedAt: time.Now(),
					})

					return s.ordersRepo.Update(orderId, taskexchange.UpdateOrderInput{
						Status:        input.Status,
						ReturnComment: input.ReturnComment,
					})
				}
			}
		}
		if order.Status == 0 {
			if order.Offer.PerformerId == user.Id && *input.Status == 1 {
				if input.SurrenderComment == nil {
					return errors.New("wrong surrender comment")
				}

				_, _ = s.eventsRepo.Create(taskexchange.Event{
					UserId:    order.Task.CustomerId,
					Message:   fmt.Sprintf("Исполнитель сдал работу по заказу #%d.", order.Id),
					Link:      "/orders/customer",
					CreatedAt: time.Now(),
				})

				return s.ordersRepo.Update(orderId, taskexchange.UpdateOrderInput{
					Status:           input.Status,
					SurrenderComment: input.SurrenderComment,
				})
			}
			if *input.Status == 3 {
				if input.CancelComment == nil {
					return errors.New("wrong cancel comment")
				}

				newTaskAmount := order.Task.Amount + 1
				err = s.tasksRepo.Update(order.Task.Id, taskexchange.UpdateTaskInput{
					Amount: &newTaskAmount,
				})
				if err != nil {
					return err
				}

				if order.Task.CustomerId == user.Id {
					_, _ = s.eventsRepo.Create(taskexchange.Event{
						UserId:    order.Offer.PerformerId,
						Message:   fmt.Sprintf("Заказчик отменил заказ #%d.", order.Id),
						Link:      "/orders/performer",
						CreatedAt: time.Now(),
					})
				} else {
					_, _ = s.eventsRepo.Create(taskexchange.Event{
						UserId:    order.Task.CustomerId,
						Message:   fmt.Sprintf("Исполнитель отменил заказ #%d.", order.Id),
						Link:      "/orders/customer",
						CreatedAt: time.Now(),
					})
				}

				return s.ordersRepo.Update(orderId, taskexchange.UpdateOrderInput{
					Status:         input.Status,
					CancelComment:  input.CancelComment,
					CanceledUserId: &user.Id,
				})
			}
		}
	}

	return nil
}

func (s *OrdersService) CountAllActive() (int, error) {
	return s.ordersRepo.CountAllActive()
}

func (s *OrdersService) GetAllCompleted() ([]taskexchange.Order, error) {
	return s.ordersRepo.GetAllCompleted()
}
