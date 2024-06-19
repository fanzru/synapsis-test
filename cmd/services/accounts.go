package services

import (
	accountshandler "synapsis-test/app/account/http"
	accountsrepo "synapsis-test/app/account/repo"
	accountsusecase "synapsis-test/app/account/usecase"
	"synapsis-test/infrastructure/config"
	"synapsis-test/infrastructure/database"
)

func RegisterServiceAccounts(db database.Connection, cfg config.Config) accountshandler.AccountHandler {
	accountsDB := accountsrepo.New(&accountsrepo.AccountsRepo{
		MySQL: db,
		Cfg:   cfg,
	})

	accountsApp := accountsusecase.New(&accountsusecase.AccountsApp{
		AccountsRepo: accountsDB,
		Cfg:          cfg,
	})

	accountHandler := accountshandler.AccountHandler{
		AccountsApp: accountsApp,
		Cfg:         cfg,
	}

	return accountHandler
}
