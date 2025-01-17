package abstract

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationQuery model.
type PaginationQuery struct {
	Limit       int
	CurrentPage int
}

// PaginatedResponse model for general responsing paginated models.
type PaginatedResponse[T any] struct {
	Items              []T `json:"items"`
	Limit              int `json:"limit"`
	CurrentPage        int `json:"current_page"`
	TotalPages         int `json:"total_pages"`
	ItemsInCurrentPage int `json:"items_in_current_page"`
	TotalItems         int `json:"total_items"`
}

func (p *PaginationQuery) SetLimit(limit string) error {
	if limit == "" {
		p.Limit = 10
		return nil
	}

	n, err := strconv.Atoi(limit)
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}
	p.Limit = n

	return nil
}

func (p *PaginationQuery) SetPage(currentPage string) error {
	if currentPage == "" {
		p.CurrentPage = 1
		return nil
	}

	n, err := strconv.Atoi(currentPage)
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}
	p.CurrentPage = n

	return nil
}

func GetPaginationFromGinCtx(c *gin.Context) (PaginationQuery, error) {
	pq := PaginationQuery{}

	if err := pq.SetPage(c.Query("current-page")); err != nil {
		return pq, fmt.Errorf("err: setting current-page: %w", err)
	}

	if err := pq.SetLimit(c.Query("limit")); err != nil {
		return pq, fmt.Errorf("err: setting limit: %w", err)
	}

	return pq, nil
}
