package repository

import (
	"context"

	"github.com/jumayevgadam/evernote-go/internal/connection"
	"github.com/jumayevgadam/evernote-go/internal/models/abstract"
	notebookModel "github.com/jumayevgadam/evernote-go/internal/models/notebooks"
	"github.com/jumayevgadam/evernote-go/internal/notebooks"
)

var _ notebooks.Repository = (*NotebookRepository)(nil)

type NotebookRepository struct {
	psqlDB connection.DB
}

func NewNotebookRepository(psqlDB connection.DB) *NotebookRepository {
	return &NotebookRepository{psqlDB: psqlDB}
}

// AddNotebook method inserts notebook details into repository.
func (r *NotebookRepository) AddNotebook(ctx context.Context, reqData notebookModel.RequestData) (int, error) {
	var notebookID int

	err := r.psqlDB.QueryRow(
		ctx,
		addNotebookQuery,
		reqData.UserID,
		reqData.Name,
		reqData.IsShared,
	).Scan(&notebookID)

	if err != nil {
		return 0, err
	}

	return notebookID, nil
}

func (r *NotebookRepository) CountNotebooksByUser(ctx context.Context, userID int) (int, error) {
	var count int

	err := r.psqlDB.Get(
		ctx,
		&count,
		countNotebooksByUserQuery,
		userID,
	)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *NotebookRepository) ListNotebooks(ctx context.Context, pgQuery abstract.PaginationQuery, userID int) ([]*notebookModel.NotebookData, error) {
	var notebooks []*notebookModel.NotebookData

	offset := (pgQuery.CurrentPage - 1) * pgQuery.Limit

	err := r.psqlDB.Select(
		ctx,
		&notebooks,
		listNotebooksQuery,
		userID,
		offset,
		pgQuery.Limit,
	)
	if err != nil {
		return nil, err
	}

	return notebooks, nil
}
