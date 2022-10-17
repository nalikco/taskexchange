package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
)

type OptionsService struct {
	repo repository.Options
}

func NewOptionService(repo repository.Options) *OptionsService {
	return &OptionsService{repo: repo}
}

func (s *OptionsService) Create(parentId int, option taskexchange.Option) (int, error) {
	return s.repo.Create(parentId, option)
}

func (s *OptionsService) GetAll(full bool) ([]taskexchange.Option, error) {
	return s.repo.GetAll(full)
}

func (s *OptionsService) GetCategories() ([]taskexchange.Option, error) {
	return s.repo.GetCategories()
}

func (s *OptionsService) GetById(id int, full bool) (taskexchange.Option, error) {
	return s.repo.GetById(id, full)
}

func (s *OptionsService) Update(id int, input taskexchange.UpdateOptionInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, input)
}

func (s *OptionsService) Delete(id int) error {
	option, err := s.repo.GetById(id, true)
	if err != nil {
		return err
	}

	if option.DeletedAt != nil {
		return s.repo.Restore(id)
	}

	return s.repo.Delete(id)
}
