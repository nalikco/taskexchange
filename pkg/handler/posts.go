package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskexchange"
)

type createPostInput struct {
	Categories *[]int  `json:"categories" binding:"required"`
	Status     *int    `json:"status" binding:"required"`
	Title      *string `json:"title" binding:"required"`
	Short      *string `json:"short" binding:"required"`
	Text       *string `json:"text" binding:"required"`
}

func (h *Handler) CreatePost(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return
	}

	var input createPostInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post := taskexchange.Post{
		Author: user,
		Status: *input.Status,
		Title:  *input.Title,
		Short:  *input.Short,
		Text:   *input.Text,
	}

	for _, categoryId := range *input.Categories {
		post.Categories = append(post.Categories, taskexchange.PostCategory{
			ID: categoryId,
		})
	}

	id, err := h.services.Posts.CreatePost(post)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type postResponse struct {
	Data taskexchange.Post `json:"data"`
}

type postsResponse struct {
	Data []taskexchange.Post `json:"data"`
}

func (h *Handler) GetAllPosts(c *gin.Context) {
	posts, err := h.services.Posts.GetAll(25, 0)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, postsResponse{
		Data: posts,
	})
}

func (h *Handler) GetPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Posts.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, postResponse{
		Data: post,
	})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input taskexchange.UpdatePostInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Posts.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) DeletePost(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Posts.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type createPostCategoryInput struct {
	Title *string `json:"title" binding:"required"`
}

func (h *Handler) CreateCategory(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return
	}

	var input createPostCategoryInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	category := taskexchange.PostCategory{
		Title: *input.Title,
	}

	id, err := h.services.Posts.CreateCategory(category)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type postCategoriesResponse struct {
	Data []taskexchange.PostCategory `json:"data"`
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := h.services.Posts.GetAllCategories()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, postCategoriesResponse{
		Data: categories,
	})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input taskexchange.UpdatePostCategoryInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Posts.UpdateCategory(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusForbidden, "forbidden")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Posts.DeleteCategory(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
