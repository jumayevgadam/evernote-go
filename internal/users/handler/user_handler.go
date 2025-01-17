package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumayevgadam/evernote-go/internal/models/abstract"
	userModel "github.com/jumayevgadam/evernote-go/internal/models/user"
	"github.com/jumayevgadam/evernote-go/internal/users"
	"github.com/jumayevgadam/evernote-go/pkg/httpError"
	"github.com/jumayevgadam/evernote-go/pkg/reqvalidator"
)

var _ users.Handler = (*UserHandler)(nil)

// UserHandler is a handler for managing users in handler layer.
type UserHandler struct {
	service users.Service
}

// NewUserHandler returns a new UserHandler.
func NewUserHandler(service users.Service) *UserHandler {
	return &UserHandler{service: service}
}

// SignUp method.
func (h *UserHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req userModel.SignUpReq
		
		err := reqvalidator.ReadRequest(c, &req)
		if err != nil {
			httpError.Response(c, err)
			return
		}

		userID, err := h.service.SignUp(c, req)
		if err != nil {
			log.Println(err)
			httpError.Response(c, err)
			return
		}

		c.JSON(http.StatusOK, abstract.SuccessResponse{
			Status:  "success",
			Data:    userID,
			Message: "completed successfully",
		})
	}
}

func (h *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq userModel.LoginReq

		err := reqvalidator.ReadRequest(c, &loginReq)
		if err != nil {
			httpError.Response(c, err)
			return
		}

		tokenStr, err := h.service.Login(c, loginReq)
		if err != nil {
			httpError.Response(c, err)
			return
		}

		c.JSON(http.StatusOK, abstract.SuccessResponse{
			Status:  "success",
			Data:    tokenStr,
			Message: "completed successfully",
		})
	}
}
