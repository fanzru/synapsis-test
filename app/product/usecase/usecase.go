package usecase

import (
	"context"
	"synapsis-test/app/product/domain/models"
	"synapsis-test/app/product/domain/params"
	"synapsis-test/app/product/repo"
	"synapsis-test/infrastructure/config"

	"github.com/ztrue/tracerr"
)

type Impl interface {
	FindProducts(ctx context.Context, params params.FindParams) (*[]models.Product, error)
	FindCategories(ctx context.Context, params params.FindParams) (*[]models.Category, error)
}

type ProductApp struct {
	ProductRepo repo.Impl
	Cfg         config.Config
}

func New(product *ProductApp) Impl {
	return product
}

func (i *ProductApp) FindProducts(ctx context.Context, params params.FindParams) (*[]models.Product, error) {
	data, err := i.ProductRepo.FindProducts(ctx, params)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	return data, nil
}

func (i *ProductApp) FindCategories(ctx context.Context, params params.FindParams) (*[]models.Category, error) {
	data, err := i.ProductRepo.FindCategory(ctx, params)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	return data, nil
}
