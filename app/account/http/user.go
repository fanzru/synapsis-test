package http

import (
	"synapsis-test/app/account/domain/request"
	"synapsis-test/pkg/contextutil"
	"synapsis-test/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h AccountHandler) LoginUser(c echo.Context) error {
	bodyRequest := &request.LoginBodyRequest{}

	err := c.Bind(bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}
	err = validator.New().Struct(bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}

	ctx := c.Request().Context()
	resp, err := h.AccountsApp.LoginUser(ctx, bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}

	return response.ResponseSuccessCreated(c, resp)
}

func (h AccountHandler) ChangePassword(c echo.Context) error {
	ctx := contextutil.OverrideContext(c)
	bodyRequest := &request.UpdatePasswordBodyRequest{}

	err := c.Bind(bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}
	err = validator.New().Struct(bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}

	err = h.AccountsApp.UpdateUser(ctx, bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}

	return response.ResponseSuccessCreated(c, nil)
}

func (h AccountHandler) CreateNewAccount(c echo.Context) error {
	ctx := contextutil.OverrideContext(c)
	bodyRequest := &request.CreateAccountBodyRequest{}

	err := c.Bind(bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}
	err = validator.New().Struct(bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}

	err = h.AccountsApp.CreateNewAccount(ctx, bodyRequest)
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}

	return response.ResponseSuccessCreated(c, nil)
}
