package notebooks

import (
	"context"

	"github.com/jumayevgadam/evernote-go/internal/models/abstract"
	notebookModel "github.com/jumayevgadam/evernote-go/internal/models/notebooks"
)

// Repository interface for notebooks.
type Repository interface {
	AddNotebook(ctx context.Context, reqData notebookModel.RequestData) (int, error)
	CountNotebooksByUser(ctx context.Context, userID int) (int, error)
	ListNotebooks(ctx context.Context, pgQueryData abstract.PaginationQuery, userID int) ([]*notebookModel.NotebookData, error)
}
