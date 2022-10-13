package service

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"taskexchange"
	"taskexchange/pkg/repository"
	"time"
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

func (s *TasksService) CreateFromExcelFile(userId int, filename string) error {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			logrus.Error(err)
		}

		if err := os.Remove(filename); err != nil {
			logrus.Error(err)
		}
	}()

	customer, err := s.usersRepo.GetById(userId, true)
	if err != nil {
		return err
	}

	sheetMap := f.GetSheetMap()

	var tasks []taskexchange.Task
	var price float64 = 0
	for _, sheetName := range sheetMap {
		task := taskexchange.Task{}

		rows, err := f.GetRows(sheetName)
		if err != nil {
			return err
		}

		if len(rows) < 2 || len(rows[0]) < 4 || len(rows[1]) < 1 {
			return errors.New("wrong excel format, check documentation")
		}

		task.Amount, err = strconv.Atoi(rows[0][3])
		if err != nil {
			return err
		}

		task.CustomerId = userId
		task.Status = 1
		task.Link = rows[0][0]
		task.Description = rows[0][1]
		task.DeliveryDate, err = time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", rows[0][2]))
		if err != nil {
			return err
		}

		mainOption, err := s.optionsRepo.GetByTitle(rows[1][0], 0)
		if err != nil {
			return err
		}

		task.Options = append(task.Options, mainOption)

		if len(rows) > 2 {
			for _, optionTitle := range rows[2] {
				option, err := s.optionsRepo.GetByTitle(optionTitle, mainOption.Id)
				if err != nil {
					return err
				}

				task.Options = append(task.Options, option)
			}
		}

		price += task.CalculatePrice()
		tasks = append(tasks, task)
	}

	if customer.Balance < price {
		return errors.New("wrong user balance")
	}
	userNewBalance := customer.Balance - price

	for _, task := range tasks {
		taskId, err := s.tasksRepo.Create(task)
		if err != nil {
			return err
		}

		for _, option := range task.Options {
			_, err = s.taskOptionsRepo.Create(taskId, option.Id)
			if err != nil {
				return err
			}
		}
	}

	updateUserInput := taskexchange.UpdateUserInput{
		Balance: &userNewBalance,
	}

	err = s.usersRepo.Update(customer.Id, updateUserInput)

	return nil
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
		customer.Balance += task.CalculatePrice()

		for _, option := range task.Options {
			taskOptionId := 0

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

	for i, task := range tasks {
		taskOptions, err := s.taskOptionsRepo.GetByTaskId(task.Id)
		if err != nil {
			return []taskexchange.Task{}, pagination, err
		}
		var taskOptionsIds []int

		for _, taskOption := range taskOptions {
			taskOptionsIds = append(taskOptionsIds, taskOption.OptionId)
		}

		options, err := s.optionsRepo.GetByIds(taskOptionsIds)
		if err != nil {
			return []taskexchange.Task{}, pagination, err
		}

		tasks[i].Options = options
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
