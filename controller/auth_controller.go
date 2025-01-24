package controller

import (
	"go-api/auth"
	"go-api/request"
	"go-api/response"
	"go-api/service"
	"go-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service *service.UserService
}

func NewAuthController(service *service.UserService) *AuthController {
	return &AuthController{Service: service}
}

func (c *AuthController) LoginHandler(ctx echo.Context) error {
	var loginRequest request.LoginRequest
	if err := ctx.Bind(&loginRequest); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, err.Error())
	}
	user, err := c.Service.GetOneUserByNameAndPassword(loginRequest.Name, loginRequest.Password)
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "User Not Found")
	}
	token, err := auth.GenerateToken(uint(user.ID))
	if err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Generate Token Failed")
	}
	var response response.LoginResponse
	response.Token = token
	return utils.Respond(ctx, http.StatusOK, response, "Success")

}
