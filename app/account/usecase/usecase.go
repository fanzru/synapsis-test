package usecase

import (
	"context"
	requests "synapsis-test/app/account/domain/request"
	"synapsis-test/app/account/domain/response"
	"synapsis-test/app/account/repo"
	"synapsis-test/infrastructure/config"
)

type Impl interface {
	CreateNewAccount(ctx context.Context, params *requests.CreateAccountBodyRequest) error
	LoginUser(ctx context.Context, params *requests.LoginBodyRequest) (*response.LoginResponse, error)
	UpdateUser(ctx context.Context, params *requests.UpdatePasswordBodyRequest) error
}

type AccountsApp struct {
	AccountsRepo repo.Impl
	Cfg          config.Config
}

func New(accounts *AccountsApp) Impl {
	return accounts
}
