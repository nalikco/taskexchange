package service

import (
	"database/sql"
	"errors"
	"taskexchange"
	"taskexchange/pkg/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUser(user taskexchange.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	if user.Type != 1 && user.Type != 2 && user.Type != 3 {
		return 0, errors.New("not valid type")
	}

	userByEmail, err := s.repo.GetByEmail(user.Email)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if userByEmail.Email == user.Email {
		return 0, errors.New("email is already taken")
	}

	return s.repo.Create(user)
}

func (s *UsersService) GetAll(full bool) ([]taskexchange.User, error) {
	return s.repo.GetAll(full)
}

func (s *UsersService) CountAll(sort taskexchange.SortUsersCount) (int, error) {
	return s.repo.CountAll(sort)
}

func (s *UsersService) GetById(id int, full bool) (taskexchange.User, error) {
	return s.repo.GetById(id, full)
}

func (s *UsersService) Update(id int, input taskexchange.UpdateUserInput) error {
	if input.Email != nil {
		userByEmail, err := s.repo.GetByEmail(*input.Email)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if userByEmail.Email == *input.Email {
			return errors.New("email is already taken")
		}
	}

	if input.Password != nil {
		*input.Password = generatePasswordHash(*input.Password)
	}

	return s.repo.Update(id, input)
}

func (s *UsersService) Delete(id int) error {
	user, err := s.repo.GetById(id, true)
	if err != nil {
		return err
	}

	if user.DeletedAt != nil {
		return s.repo.Restore(id)
	}

	return s.repo.Delete(id)
}
