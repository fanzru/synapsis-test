package services

import (
	producthandler "synapsis-test/app/product/http"
	productrepo "synapsis-test/app/product/repo"
	productusecase "synapsis-test/app/product/usecase"
	"synapsis-test/infrastructure/config"
	"synapsis-test/infrastructure/database"
)

func RegisterServiceProduct(db database.Connection, cfg config.Config) producthandler.ProductHandler {
	productDB := productrepo.New(&productrepo.ProductRepo{
		MySQL: db,
		Cfg:   cfg,
	})

	productUsecase := productusecase.New(&productusecase.ProductApp{
		ProductRepo: productDB,
		Cfg:         cfg,
	})

	productHandler := producthandler.ProductHandler{
		ProductUsecase: productUsecase,
		Cfg:            cfg,
	}

	return productHandler
}
