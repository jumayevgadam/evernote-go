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
// @Summary User Register
// @Tags users
// @Description user register to evernote
// @Accept json
// @Produce json
// @Param input body userModel.SignUpReq true "sign up info"
// @Success 200 {object} abstract.SuccessResponse
// @Failure 400 {object} httpError.RestErr
// @Failure 500 {object} httpError.RestErr
// @Failure default {object} httpError.RestErr
// @Router /auth/register [post]
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

// Login func for users.
// @Summary User Login
// @Tags users
// @Description user login to evernote
// @Accept json
// @Produce json
// @Param input body userModel.LoginReq true "login info"
// @Success 200 {object} abstract.SuccessResponse
// @Failure 400 {object} httpError.RestErr
// @Failure 500 {object} httpError.RestErr
// @Failure default {object} httpError.RestErr
// @Router /auth/login [post]
func (h *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq userModel.LoginReq

		err := reqvalidator.ReadRequest(c, &loginReq)
		if err != nil {
			httpError.Response(c, err)
			return
		}

		tokens, err := h.service.Login(c, loginReq)
		if err != nil {
			httpError.Response(c, err)
			return
		}

		c.JSON(http.StatusOK, abstract.SuccessResponse{
			Status:  "success",
			Data:    tokens,
			Message: "completed successfully",
		})
	}
}
