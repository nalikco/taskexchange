package service

import (
	"errors"
	"fmt"
	"taskexchange"
	"taskexchange/pkg/repository"
)

type TasksService struct {
	tasksRepo       repository.Tasks
	taskOptionsRepo repository.TaskOptions
	optionsRepo     repository.Options
	usersRepo       repository.Users
}

func NewTasksService(tasksRepo repository.Tasks, taskOptionsRepo repository.TaskOptions, usersRepo repository.Users, optionsRepo repository.Options) *TasksService {
	return &TasksService{
		tasksRepo:       tasksRepo,
		taskOptionsRepo: taskOptionsRepo,
		optionsRepo:     optionsRepo,
		usersRepo:       usersRepo,
	}
}

func (s *TasksService) Create(task taskexchange.Task) (int, error) {
	customer, err := s.usersRepo.GetById(task.CustomerId, true)
	if err != nil {
		return 0, err
	}

	taskPrice := task.CalculatePrice()

	if customer.Balance < taskPrice {
		return 0, errors.New("wrong user balance")
	}
	userNewBalance := customer.Balance - taskPrice

	taskId, err := s.tasksRepo.Create(task)
	if err != nil {
		return 0, err
	}

	for _, option := range task.Options {
		_, err = s.taskOptionsRepo.Create(taskId, option.Id)
		if err != nil {
			return 0, err
		}
	}

	updateUserInput := taskexchange.UpdateUserInput{
		Balance: &userNewBalance,
	}

	err = s.usersRepo.Update(customer.Id, updateUserInput)

	return taskId, nil
}

func (s *TasksService) Update(id int, input taskexchange.UpdateTaskInput) error {
	if input.Options != nil {
		task, err := s.GetById(id)
		if err != nil {
			return err
		}

		taskOptions, err := s.taskOptionsRepo.GetByTaskId(task.Id)
		if err != nil {
			return err
		}

		customer, err := s.usersRepo.GetById(task.CustomerId, true)
		if err != nil {
			return err
		}

		for _, option := range task.Options {
			taskOptionId := 0

			customer.Balance += option.Price
			for _, taskOption := range taskOptions {
				if taskOption.OptionId == option.Id {
					taskOptionId = taskOption.Id
				}
			}

			if taskOptionId != 0 {
				err = s.taskOptionsRepo.Delete(taskOptionId)
				if err != nil {
					return err
				}
			}
		}

		task.Options = []taskexchange.Option{}

		for _, optionId := range *input.Options {
			option, err := s.optionsRepo.GetById(optionId)
			if err != nil {
				return errors.New(fmt.Sprintf("wrong option id: %d", optionId))
			}

			if option.ParentId != nil {
				var parentIdFound = false

				for _, parentId := range *input.Options {
					if parentId == *option.ParentId {
						parentIdFound = true
					}
				}

				if !parentIdFound {
					return errors.New(fmt.Sprintf("parent id %d not found in options array for option: %d", *option.ParentId, optionId))
				}
			}

			task.Options = append(task.Options, option)
		}

		taskPrice := task.CalculatePrice()

		if customer.Balance < taskPrice {
			return errors.New("wrong user balance")
		}
		customer.Balance -= taskPrice

		for _, option := range task.Options {
			_, err := s.taskOptionsRepo.Create(id, option.Id)

			if err != nil {
				return err
			}
		}

		updateUserInput := taskexchange.UpdateUserInput{
			Balance: &customer.Balance,
		}

		err = s.usersRepo.Update(customer.Id, updateUserInput)
	}

	return s.tasksRepo.Update(id, input)
}

func (s *TasksService) GetById(id int) (taskexchange.Task, error) {
	task, err := s.tasksRepo.GetById(id)
	if err != nil {
		return taskexchange.Task{}, err
	}

	taskOptions, err := s.taskOptionsRepo.GetByTaskId(task.Id)
	if err != nil {
		return taskexchange.Task{}, err
	}
	var taskOptionsIds []int

	for _, taskOption := range taskOptions {
		taskOptionsIds = append(taskOptionsIds, taskOption.OptionId)
	}

	options, err := s.optionsRepo.GetByIds(taskOptionsIds)
	if err != nil {
		return taskexchange.Task{}, err
	}

	task.Options = options

	return task, nil
}

func (s *TasksService) GetAll(userId int, pagination taskexchange.Pagination) ([]taskexchange.Task, taskexchange.Pagination, error) {
	var count int
	var err error

	if userId == 0 {
		count, err = s.tasksRepo.CountAll()
	} else {
		count, err = s.tasksRepo.CountAllByUser(userId)
	}

	if err != nil {
		return []taskexchange.Task{}, pagination, err
	}

	pagination.Calculate(count)

	var tasks []taskexchange.Task

	if userId == 0 {
		tasks, err = s.tasksRepo.FindAll(pagination.Limit, pagination.Offset)
	} else {
		tasks, err = s.tasksRepo.FindAllByUser(userId, pagination.Limit, pagination.Offset)
	}

	return tasks, pagination, err
}

func (s *TasksService) Delete(id int, task taskexchange.Task, customerId int) error {
	customer, err := s.usersRepo.GetById(customerId, true)
	if err != nil {
		return err
	}

	newCustomerBalance := customer.Balance + task.CalculatePrice()

	err = s.usersRepo.Update(customer.Id, taskexchange.UpdateUserInput{
		Balance: &newCustomerBalance,
	})
	if err != nil {
		return err
	}

	return s.tasksRepo.Delete(id)
}
