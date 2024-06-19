package http

import (
	"strconv"
	"synapsis-test/app/product/domain/params"
	"synapsis-test/app/product/usecase"
	"synapsis-test/infrastructure/config"
	"synapsis-test/pkg/response"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUsecase usecase.Impl
	Cfg            config.Config
}

func (h *ProductHandler) FindProducts(c echo.Context) error {
	categoryId := c.QueryParam("categoryId")
	categoryIdNumber, _ := strconv.Atoi(categoryId)

	ctx := c.Request().Context()
	resp, err := h.ProductUsecase.FindProducts(ctx, params.FindParams{
		CategoryID: int64(categoryIdNumber),
	})
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}

	return response.ResponseSuccessCreated(c, resp)
}

func (h *ProductHandler) FindCategories(c echo.Context) error {
	categoryId := c.QueryParam("categoryId")
	categoryIdNumber, _ := strconv.Atoi(categoryId)

	ctx := c.Request().Context()
	resp, err := h.ProductUsecase.FindCategories(ctx, params.FindParams{
		CategoryID: int64(categoryIdNumber),
	})
	if err != nil {
		return response.ResponseErrorBadRequest(c, err)
	}

	return response.ResponseSuccessCreated(c, resp)
}
