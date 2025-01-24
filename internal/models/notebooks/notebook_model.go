package notebooks

import "time"

// Request model for creating notebook.
type Request struct {
	Name   string `json:"notebook_name" validate:"required"`
	UserID int    `json:"-"`
}

// RequestData model is db model.
type RequestData struct {
	Name     string `db:"name"`
	UserID   int    `db:"user_id"`
	IsShared bool   `db:"is_shared"`
}

// ToPsqlDBStorage func sends request model to db.
func (r *Request) ToPsqlDBStorage() RequestData {
	return RequestData{
		Name:   r.Name,
		UserID: r.UserID,
	}
}

// Notebook model.
type Notebook struct {
	Name      string    `json:"notebookName"`
	ID        int       `json:"notebookID"`
	UserID    int       `json:"userID"`
	IsShared  bool      `json:"isShared"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NotebookData model is db model.
type NotebookData struct {
	Name      string    `db:"name"`
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	IsShared  bool      `db:"is_shared"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (n *NotebookData) ToServer() *Notebook {
	return &Notebook{
		Name:      n.Name,
		ID:        n.ID,
		UserID:    n.UserID,
		IsShared:  n.IsShared,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}
