package service

import (
	"errors"
	"gorm.io/gorm"
	"taskexchange"
	"taskexchange/pkg/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUser(user taskexchange.User) (taskexchange.User, error) {
	user.Password = generatePasswordHash(user.Password)

	if user.Type != 1 && user.Type != 2 && user.Type != 3 {
		return taskexchange.User{}, errors.New("not valid type")
	}

	userByEmail, err := s.repo.GetByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return taskexchange.User{}, err
	}
	if userByEmail.Email == user.Email {
		return taskexchange.User{}, errors.New("email is already taken")
	}

	userByUsername, err := s.repo.GetByUsername(user.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return taskexchange.User{}, err
	}
	if userByUsername.Username == user.Username {
		return taskexchange.User{}, errors.New("username is already taken")
	}

	return s.repo.Create(user)
}

func (s *UsersService) GetAll() ([]taskexchange.User, error) {
	return s.repo.GetAll()
}

func (s *UsersService) GetAllHidden() ([]taskexchange.UserHidden, error) {
	return s.repo.GetAllHidden()
}

func (s *UsersService) CountAll(sort taskexchange.SortUsersCount) (int64, error) {
	return s.repo.CountAll(sort)
}

func (s *UsersService) GetById(id int) (taskexchange.User, error) {
	return s.repo.GetById(id)
}

func (s *UsersService) GetByIdHidden(id int) (taskexchange.UserHidden, error) {
	return s.repo.GetByIdHidden(id)
}

func (s *UsersService) Update(user taskexchange.User, input taskexchange.UpdateUserInput) error {
	if input.Email != nil {
		userByEmail, err := s.repo.GetByEmail(*input.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if userByEmail.Email == *input.Email {
			return errors.New("email is already taken")
		}

		user.Email = *input.Email
	}

	if input.Password != nil {
		user.Password = generatePasswordHash(*input.Password)
	}

	if input.Username != nil {
		userByUsername, err := s.repo.GetByUsername(*input.Username)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if userByUsername.Username == *input.Username {
			return errors.New("username is already taken")
		}

		user.Username = *input.Username
	}

	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		user.LastName = *input.LastName
	}

	if input.Type != nil {
		user.Type = *input.Type
	}

	if input.Balance != nil {
		user.Balance = *input.Balance
	}

	if input.Points != nil {
		user.Points = *input.Points
	}

	return s.repo.Update(user)
}

func (s *UsersService) Delete(id int) error {
	user, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	if user.DeletedAt != nil {
		return s.repo.Restore(user)
	}

	return s.repo.Delete(user)
}
