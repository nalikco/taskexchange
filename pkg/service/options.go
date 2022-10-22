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

func (s *OptionsService) Create(option taskexchange.Option) (taskexchange.Option, error) {
	return s.repo.Create(option)
}

func (s *OptionsService) GetAll(sort taskexchange.SortOptions) ([]taskexchange.Option, error) {
	return s.repo.GetAll(sort)
}

func (s *OptionsService) GetCategories() ([]taskexchange.Option, error) {
	return s.repo.GetCategories()
}

func (s *OptionsService) GetById(id int, deleted bool) (taskexchange.Option, error) {
	return s.repo.GetById(id, deleted)
}

func (s *OptionsService) Update(id int, input taskexchange.UpdateOptionInput) error {
	option, err := s.repo.GetById(id, true)
	if err != nil {
		return err
	}

	if input.ParentId != nil {
		if *input.ParentId == 0 {
			option.ParentID = nil
			option.Parent = nil
		} else {
			input.Short = nil
			option.Short = nil

			parent, err := s.repo.GetById(*input.ParentId, true)
			if err != nil {
				return err
			}

			option.Parent = &parent
		}
	}

	if input.Short != nil {
		option.Short = input.Short
	}

	if input.Title != nil {
		option.Title = *input.Title
	}

	if input.Price != nil {
		option.Price = *input.Price
	}

	return s.repo.Update(option)
}

func (s *OptionsService) Delete(id int) error {
	option, err := s.repo.GetById(id, true)
	if err != nil {
		return err
	}

	if option.DeletedAt != nil {
		return s.repo.Restore(option)
	}

	return s.repo.Delete(option)
}
