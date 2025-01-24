package notebooks

import (
	"context"

	"github.com/jumayevgadam/evernote-go/internal/models/abstract"
	notebookModel "github.com/jumayevgadam/evernote-go/internal/models/notebooks"
)

// Service interface for notebooks.
type Service interface {
	AddNotebook(ctx context.Context, req notebookModel.Request) (int, error)
	ListNotebooks(ctx context.Context, pgParams abstract.PaginationQuery, userID int) (
		abstract.PaginatedResponse[*notebookModel.Notebook], error)
}
