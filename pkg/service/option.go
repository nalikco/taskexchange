package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
)

type OptionService struct {
	repo repository.Option
}

func NewOptionService(repo repository.Option) *OptionService {
	return &OptionService{repo: repo}
}

func (s *OptionService) Create(parentId int, option taskexchange.Option) (int, error) {
	return s.repo.Create(parentId, option)
}

func (s *OptionService) GetAll() ([]taskexchange.Option, error) {
	return s.repo.GetAll()
}

func (s *OptionService) GetById(id int) (taskexchange.Option, error) {
	return s.repo.GetById(id)
}

func (s *OptionService) Update(id int, input taskexchange.UpdateOptionInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, input)
}

func (s *OptionService) Delete(id int) error {
	return s.repo.Delete(id)
}
