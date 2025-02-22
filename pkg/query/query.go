// Package query provides support for query paging.
package query

import "github.com/3bd-dev/wallet-service/pkg/page"

// Result is the data model used when returning a query result.
type Result[T any] struct {
	Items       []T   `json:"items"`
	Total       int64 `json:"total"`
	Page        int   `json:"page"`
	RowsPerPage int   `json:"rowsPerPage"`
}

// NewResult constructs a result value to return query results.
func NewResult[T any](items []T, total int64, page page.Page) Result[T] {
	return Result[T]{
		Items:       items,
		Total:       total,
		Page:        page.Number(),
		RowsPerPage: page.RowsPerPage(),
	}
}
