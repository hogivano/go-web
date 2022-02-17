package handler

import (
	"go-web/helper"
	"go-web/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) LoginUser(ctx *gin.Context) {
	var input user.LoginUserInput

	err := ctx.ShouldBindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}

	ctx.JSON(http.StatusOK, nil)
}

func (h *userHandler) RegisterUser(ctx *gin.Context) {
	var input user.RegisterUserInput

	err := ctx.ShouldBindJSON(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, input))
	}

	user, err := h.service.RegisterUser(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(err.Error(), http.StatusBadRequest, input))
	}

	ctx.JSON(http.StatusOK, helper.APIResponse("Success register account", http.StatusOK, user))
}
