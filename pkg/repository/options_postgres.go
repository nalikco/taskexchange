package repository

import (
	"gorm.io/gorm"
	"taskexchange"
	"time"
)

type OptionsPostgres struct {
	db *gorm.DB
}

func NewOptionsPostgres(db *gorm.DB) *OptionsPostgres {
	return &OptionsPostgres{db: db}
}

func (r *OptionsPostgres) Create(option taskexchange.Option) (taskexchange.Option, error) {
	result := r.db.Table("options").Create(&option)

	return option, result.Error
}

func (r *OptionsPostgres) GetAll(sort taskexchange.SortOptions) ([]taskexchange.Option, error) {
	var options []taskexchange.Option

	query := r.db.Table("options").Preload("Parent")
	if sort.Deleted == false {
		query = query.Preload("Options", "deleted_at is null").Where("deleted_at is null")
	} else {
		query = query.Preload("Options")
	}
	result := query.Order("created_at DESC").Find(&options)

	return options, result.Error
}

func (r *OptionsPostgres) GetCategories() ([]taskexchange.Option, error) {
	var options []taskexchange.Option
	result := r.db.Table("options").Where("parent_id is null").Order("created_at DESC").Find(&options)

	return options, result.Error
}

func (r *OptionsPostgres) GetById(id int, deleted bool) (taskexchange.Option, error) {
	var option taskexchange.Option

	query := r.db.Table("options").Preload("Parent")
	if !deleted {
		query = query.Preload("Options", "deleted_at is null").Where("deleted_at is null")
	} else {
		query = query.Preload("Options")
	}
	result := query.First(&option, id)

	return option, result.Error
}

func (r *OptionsPostgres) GetByTitle(title string, parentId int) (taskexchange.Option, error) {
	var option taskexchange.Option

	query := r.db.Table("options").Preload("Parent").Preload("Options").Where("title = ?", title)
	if parentId == 0 {
		query = query.Where("parent_id = ?", parentId)
	}

	result := query.First(&option)

	return option, result.Error
}

func (r *OptionsPostgres) GetByIds(ids []int) ([]taskexchange.Option, error) {
	var options []taskexchange.Option

	result := r.db.Table("options").Preload("Parent").Preload("Options").Where("id IN ?", ids).Find(&options)

	return options, result.Error
}

func (r *OptionsPostgres) Update(option taskexchange.Option) error {
	result := r.db.Table("options").Save(&option)

	return result.Error
}

func (r *OptionsPostgres) Delete(option taskexchange.Option) error {
	currentTime := time.Now()
	option.DeletedAt = &currentTime
	result := r.db.Table("options").Save(&option)

	return result.Error
}

func (r *OptionsPostgres) Restore(option taskexchange.Option) error {
	option.DeletedAt = nil
	result := r.db.Table("options").Save(&option)

	return result.Error
}
