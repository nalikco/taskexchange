package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
)

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(repo repository.Posts) *PostsService {
	return &PostsService{repo: repo}
}

func (s *PostsService) CreateCategory(category taskexchange.PostCategory) (int, error) {
	return s.repo.CreateCategory(category)
}

func (s *PostsService) CreatePost(post taskexchange.Post) (int, error) {
	return s.repo.CreatePost(post)
}

func (s *PostsService) GetById(id int) (taskexchange.Post, error) {
	return s.repo.GetById(id, false)
}

func (s *PostsService) GetCategoryById(id int) (taskexchange.PostCategory, error) {
	return s.repo.GetCategoryById(id)
}

func (s *PostsService) GetCategoriesById(ids []int) ([]taskexchange.PostCategory, error) {
	return s.repo.GetCategoriesById(ids)
}

func (s *PostsService) GetAll(limit, offset int) ([]taskexchange.Post, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *PostsService) GetAllCategories() ([]taskexchange.PostCategory, error) {
	return s.repo.GetAllCategories()
}

func (s *PostsService) Update(id int, input taskexchange.UpdatePostInput) error {
	return s.repo.Update(id, input)
}

func (s *PostsService) UpdateCategory(id int, input taskexchange.UpdatePostCategoryInput) error {
	return s.repo.UpdateCategory(id, input)
}

func (s *PostsService) Delete(id int) error {
	post, err := s.repo.GetById(id, true)
	if err != nil {
		return err
	}
	if post.DeletedAt != nil {
		return s.repo.Restore(post.ID)
	}

	return s.repo.Delete(post.ID)
}

func (s *PostsService) DeleteCategory(id int) error {
	category, err := s.repo.GetCategoryById(id)
	if err != nil {
		return err
	}
	if category.DeletedAt != nil {
		return s.repo.RestoreCategory(category.ID)
	}

	return s.repo.DeleteCategory(category.ID)
}
