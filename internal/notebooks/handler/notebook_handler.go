package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/middlewares"
	"github.com/jumayevgadam/evernote-go/internal/notebooks"
	"github.com/jumayevgadam/evernote-go/pkg/httpError"
	"github.com/jumayevgadam/evernote-go/pkg/reqvalidator"

	"github.com/jumayevgadam/evernote-go/internal/models/abstract"
	notebookModel "github.com/jumayevgadam/evernote-go/internal/models/notebooks"
)

var _ notebooks.Handler = (*NoteBookHandler)(nil)

// NoteBook handler for notebooks.
type NoteBookHandler struct {
	service notebooks.Service
}

func NewNotebookHandler(service notebooks.Service) *NoteBookHandler {
	return &NoteBookHandler{service: service}
}

// AddNotebook handler.
// @Summary Add Notebook
// @Tags notebooks
// @Description creating a new notebook
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body notebookModel.Request true "notebook request info"
// @Success 200 {object} abstract.SuccessResponse
// @Failure 400 {object} httpError.RestErr
// @Failure 500 {object} httpError.RestErr
// @Failure default {object} httpError.RestErr
// @Router /notebooks [post]
func (h *NoteBookHandler) AddNotebook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req notebookModel.Request

		userID, err := middlewares.GetUserIDFromCtx(c)
		if err != nil {
			httpError.Response(c, err)
			return
		}
		req.UserID = userID

		err = reqvalidator.ReadRequest(c, &req)
		if err != nil {
			httpError.Response(c, err)
			return
		}

		notebookID, err := h.service.AddNotebook(c, req)
		if err != nil {
			httpError.Response(c, err)
			return
		}

		c.JSON(http.StatusOK, abstract.SuccessResponse{
			Status:  "success",
			Data:    notebookID,
			Message: "completed successfully",
		})
	}
}

// ListNotebooks handler.
// @Summary List Notebooks
// @Tags notebooks
// @Description listing all notebooks.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param current-page int false "page number" Format(page)
// @Param limit query int false "number of elements per page" Format(limit)
// @Success 200 {object} []*notebookModel.NotebookData
// @Failure 400 {object} httpError.RestErr
// @Failure 500 {object} httpError.RestErr
// @Failure default {object} httpError.RestErr
// @Router /notebooks [get]
func (h *NoteBookHandler) ListNotebooks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get user id from gin context.
		userID, err := middlewares.GetUserIDFromCtx(ctx)
		if err != nil {
			httpError.Response(ctx, err)
			return
		}

		pq, err := abstract.GetPaginationFromGinCtx(ctx)
		if err != nil {
			httpError.Response(ctx, err)
			return
		}

		res, err := h.service.ListNotebooks(ctx, pq, userID)
		if err != nil {
			httpError.Response(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
