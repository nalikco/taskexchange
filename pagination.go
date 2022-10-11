package taskexchange

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	CurrentPage int
	PerPage     int
	Limit       int
	Offset      int
	Pages       int
	Count       int
}

func NewPagination(c *gin.Context, defaultPage, defaultPerPage int) Pagination {
	var page int
	pageParam, ok := c.GetQuery("page")
	if !ok || pageParam == "" {
		page = 0
	} else {
		page, _ = strconv.Atoi(pageParam)
	}

	var perPage int
	perPageParam, ok := c.GetQuery("per_page")
	if !ok || perPageParam == "" {
		perPage = 0
	} else {
		perPage, _ = strconv.Atoi(perPageParam)
	}

	return Pagination{
		CurrentPage: page,
		PerPage:     perPage,
	}
}

func (p *Pagination) Calculate(count int) {
	p.Pages = int(math.Ceil(float64(count) / float64(p.PerPage)))
	p.Offset = (p.PerPage * p.CurrentPage) - p.PerPage
	p.Limit = p.PerPage
	p.Count = count
}
