package notebooks

// Request model for creating notebook.
type Request struct {
	UserID int    `json:"-"`
	Name   string `json:"notebook_name" validate:"required"`
}

// RequestData model is db model.
type RequestData struct {
	IsShared bool   `db:"is_shared"`
	UserID   int    `db:"user_id"`
	Name     string `db:"name"`
}

// ToPsqlDBStorage func sends request model to db.
func (r *Request) ToPsqlDBStorage() RequestData {
	return RequestData{
		UserID: r.UserID,
		Name:   r.Name,
	}
}
