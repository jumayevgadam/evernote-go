package service

import (
	"context"

	"github.com/jumayevgadam/evernote-go/internal/database"
	"github.com/jumayevgadam/evernote-go/internal/notebooks"
	"github.com/samber/lo"

	"github.com/jumayevgadam/evernote-go/internal/models/abstract"
	notebookModel "github.com/jumayevgadam/evernote-go/internal/models/notebooks"
)

var _ notebooks.Service = (*NotebookService)(nil)

type NotebookService struct {
	repo database.DataStore
}

func NewNotebookService(repo database.DataStore) *NotebookService {
	return &NotebookService{repo: repo}
}

// AddNotebook service.
func (s *NotebookService) AddNotebook(ctx context.Context, req notebookModel.Request) (int, error) {
	notebookID, err := s.repo.NotebooksRepo().AddNotebook(ctx, req.ToPsqlDBStorage())
	if err != nil {
		return 0, err
	}

	return notebookID, nil
}

// ListNotebook service.
func (s *NotebookService) ListNotebooks(ctx context.Context, pgQuery abstract.PaginationQuery, userID int) (
	abstract.PaginatedResponse[*notebookModel.Notebook], error,
) {
	// declare variables.
	var (
		allNotebookData      []*notebookModel.NotebookData
		notebookListResponse abstract.PaginatedResponse[*notebookModel.Notebook]
		err                  error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		var notebookCount int

		notebookCount, err = db.NotebooksRepo().CountNotebooksByUser(ctx, userID)
		if err != nil {
			return err
		}

		notebookListResponse.TotalItems = notebookCount

		allNotebookData, err = db.NotebooksRepo().ListNotebooks(ctx, pgQuery, userID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return abstract.PaginatedResponse[*notebookModel.Notebook]{}, err
	}

	notebookList := lo.Map(
		allNotebookData,
		func(item *notebookModel.NotebookData, _ int) *notebookModel.Notebook {
			return item.ToServer()
		},
	)

	notebookListResponse.Items = notebookList
	notebookListResponse.CurrentPage = pgQuery.CurrentPage
	notebookListResponse.Limit = pgQuery.Limit
	notebookListResponse.ItemsInCurrentPage = len(notebookList)

	return notebookListResponse, nil
}
