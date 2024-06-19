package http

import (
	accountsapp "synapsis-test/app/account/usecase"
	"synapsis-test/infrastructure/config"
)

type AccountHandler struct {
	AccountsApp accountsapp.Impl
	Cfg         config.Config
}
