package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"taskexchange"
	"time"
)

type PostsPostgres struct {
	db *sqlx.DB
}

type postDb struct {
	ID        int        `db:"id"`
	AuthorId  int        `db:"user_id"`
	MainImage string     `db:"main_image"`
	Status    int        `db:"status"`
	Title     string     `db:"title"`
	Short     string     `db:"short"`
	Text      string     `db:"text"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type postCategoryDb struct {
	ID         int `db:"id"`
	CategoryId int `db:"category_id"`
	PostId     int `db:"post_id"`
}

func NewPostsPostgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{
		db: db,
	}
}

func (r *PostsPostgres) CreateCategory(category taskexchange.PostCategory) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", postCategoriesTable)
	row := r.db.QueryRow(query, category.Title)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostsPostgres) CreatePost(post taskexchange.Post) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, status, title, short, text, main_image) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", postsTable)
	row := r.db.QueryRow(query, post.Author.Id, post.Status, post.Title, post.Short, post.Text, post.MainImage)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	for _, postCategory := range post.Categories {
		query = fmt.Sprintf("INSERT INTO %s (category_id, post_id) VALUES ($1, $2) RETURNING id", postCategoryTable)
		if _, err := r.db.Exec(query, postCategory.ID, id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (r *PostsPostgres) GetById(id int, deleted bool) (taskexchange.Post, error) {
	var post taskexchange.Post
	var postFromDb postDb
	var query string

	if deleted {
		query = fmt.Sprintf("SELECT * FROM %s WHERE id=$1", postsTable)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null AND id=$1", postsTable)
	}

	err := r.db.Get(&postFromDb, query, id)
	if err != nil {
		return post, err
	}

	post = taskexchange.Post{
		ID:        postFromDb.ID,
		Status:    postFromDb.Status,
		MainImage: postFromDb.MainImage,
		Title:     postFromDb.Title,
		Short:     postFromDb.Short,
		Text:      postFromDb.Text,
		CreatedAt: postFromDb.CreatedAt,
		DeletedAt: postFromDb.DeletedAt,
	}

	query = fmt.Sprintf("SELECT id,username,first_name,last_name,type,last_online,created_at FROM %s WHERE id=$1", usersTable)
	err = r.db.Get(&post.Author, query, postFromDb.AuthorId)
	if err != nil {
		return post, err
	}

	var postCategoriesDb []postCategoryDb
	query = fmt.Sprintf("SELECT * FROM %s WHERE post_id=$1", postCategoryTable)
	err = r.db.Select(&postCategoriesDb, query, postFromDb.ID)
	if err != nil {
		return post, err
	}

	var categoryIds []string
	for _, categoryDb := range postCategoriesDb {
		categoryIds = append(categoryIds, strconv.Itoa(categoryDb.CategoryId))
	}

	if len(categoryIds) > 0 {
		categoryIdsString := strings.Join(categoryIds, ", ")

		query = fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)", postCategoriesTable, categoryIdsString)
		err = r.db.Select(&post.Categories, query)
		if err != nil {
			return post, err
		}
	}

	return post, nil
}

func (r *PostsPostgres) GetCategoryById(id int) (taskexchange.PostCategory, error) {
	var category taskexchange.PostCategory

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", postCategoriesTable)
	err := r.db.Get(&category, query, id)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *PostsPostgres) GetCategoriesById(ids []int) ([]taskexchange.PostCategory, error) {
	var categories []taskexchange.PostCategory

	var categoryIds []string
	for _, id := range ids {
		categoryIds = append(categoryIds, strconv.Itoa(id))
	}

	categoryIdsString := strings.Join(categoryIds, ", ")

	query := fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)", postCategoriesTable, categoryIdsString)
	err := r.db.Select(&categories, query)
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *PostsPostgres) GetAll(limit, offset int) ([]taskexchange.Post, error) {
	var posts []taskexchange.Post
	var postsDb []postDb

	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null ORDER BY created_at DESC LIMIT %d OFFSET %d", postsTable, limit, offset)
	err := r.db.Select(&postsDb, query)
	if err != nil {
		return posts, err
	}

	for _, db := range postsDb {
		post := taskexchange.Post{
			ID:        db.ID,
			Status:    db.Status,
			MainImage: db.MainImage,
			Title:     db.Title,
			Short:     db.Short,
			Text:      db.Text,
			CreatedAt: db.CreatedAt,
			DeletedAt: db.DeletedAt,
		}
		var postCategoriesDb []postCategoryDb

		query = fmt.Sprintf("SELECT id,username,first_name,last_name,type,last_online,created_at FROM %s WHERE id=$1", usersTable)
		err = r.db.Get(&post.Author, query, db.AuthorId)
		if err != nil {
			return posts, err
		}

		query = fmt.Sprintf("SELECT * FROM %s WHERE post_id=$1", postCategoryTable)
		err = r.db.Select(&postCategoriesDb, query, post.ID)
		if err != nil {
			return posts, err
		}

		var categoryIds []string
		for _, categoryDb := range postCategoriesDb {
			categoryIds = append(categoryIds, strconv.Itoa(categoryDb.CategoryId))
		}

		if len(categoryIds) > 0 {
			categoryIdsString := strings.Join(categoryIds, ", ")

			query = fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)", postCategoriesTable, categoryIdsString)
			err = r.db.Select(&post.Categories, query)
			if err != nil {
				return posts, err
			}
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostsPostgres) GetAllCategories() ([]taskexchange.PostCategory, error) {
	var categories []taskexchange.PostCategory

	query := fmt.Sprintf("SELECT * FROM %s WHERE deleted_at is null ORDER BY created_at DESC", postCategoriesTable)
	err := r.db.Select(&categories, query)
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *PostsPostgres) Update(id int, input taskexchange.UpdatePostInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.MainImage != nil {
		setValues = append(setValues, fmt.Sprintf("main_image=$%d", argId))
		args = append(args, *input.MainImage)
		argId++
	}

	if input.Short != nil {
		setValues = append(setValues, fmt.Sprintf("short=$%d", argId))
		args = append(args, *input.Short)
		argId++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, *input.Text)
		argId++
	}

	args = append(args, id)

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", postsTable, setQuery, argId)
	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	if input.Categories != nil {
		query = fmt.Sprintf("DELETE FROM %s WHERE post_id=$1", postCategoryTable)
		if _, err := r.db.Exec(query, id); err != nil {
			return err
		}

		for _, postCategory := range *input.Categories {
			query = fmt.Sprintf("INSERT INTO %s (category_id, post_id) VALUES ($1, $2) RETURNING id", postCategoryTable)
			if _, err := r.db.Exec(query, postCategory, id); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *PostsPostgres) UpdateCategory(id int, input taskexchange.UpdatePostCategoryInput) error {
	if input.Title != nil {
		query := fmt.Sprintf("UPDATE %s SET title=$2 WHERE id=$1", postCategoriesTable)
		_, err := r.db.Exec(query, id, input.Title)

		return err
	}
	return nil
}

func (r *PostsPostgres) Delete(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=now() WHERE id=$1", postsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *PostsPostgres) DeleteCategory(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=now() WHERE id=$1", postCategoriesTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *PostsPostgres) Restore(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=null WHERE id=$1", postsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *PostsPostgres) RestoreCategory(id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=null WHERE id=$1", postCategoriesTable)
	_, err := r.db.Exec(query, id)

	return err
}
