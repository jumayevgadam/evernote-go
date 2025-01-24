package postgres

import (
	"sync"

	"github.com/jumayevgadam/evernote-go/internal/connection"
	"github.com/jumayevgadam/evernote-go/internal/database"
	"github.com/jumayevgadam/evernote-go/internal/notebooks"
	notebookRepository "github.com/jumayevgadam/evernote-go/internal/notebooks/repository"
	"github.com/jumayevgadam/evernote-go/internal/users"
	userRepository "github.com/jumayevgadam/evernote-go/internal/users/repository"
)

var _ database.DataStore = (*DataStore)(nil)

// DataStore is a postgres implementation of database.DataStore.
type DataStore struct {
	db           connection.DB
	user         users.Repository
	userInit     sync.Once
	notebook     notebooks.Repository
	notebookInit sync.Once
}

// NewDataStore returns a new DataStore.
func NewDataStore(db connection.DBOps) *DataStore {
	return &DataStore{
		db: db,
	}
}

// UsersRepo returns a user repository.
func (d *DataStore) UsersRepo() users.Repository {
	d.userInit.Do(func() {
		d.user = userRepository.NewUserRepository(d.db)
	})

	return d.user
}

// NotebooksRepo returns a notebook repository.
func (d *DataStore) NotebooksRepo() notebooks.Repository {
	d.notebookInit.Do(func() {
		d.notebook = notebookRepository.NewNotebookRepository(d.db)
	})

	return d.notebook
}
