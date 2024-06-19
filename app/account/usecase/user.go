package usecase

import (
	"context"
	"synapsis-test/app/account/domain/constant"
	"synapsis-test/app/account/domain/models"
	"synapsis-test/app/account/domain/param"
	requests "synapsis-test/app/account/domain/request"
	"synapsis-test/app/account/domain/response"
	"synapsis-test/pkg/contextutil"
	"synapsis-test/pkg/jwt"
	"time"

	"github.com/ztrue/tracerr"
	"golang.org/x/crypto/bcrypt"
)

func (a *AccountsApp) LoginUser(ctx context.Context, params *requests.LoginBodyRequest) (*response.LoginResponse, error) {
	user := &models.User{}

	err := a.AccountsRepo.FindUser(ctx, user, param.FindParam{
		Username: params.Username,
		Status:   string(constant.Actived),
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, tracerr.Wrap(constant.ErrWrongPassword)
	}

	token, err := jwt.EncodeToken(*user, a.Cfg.JWTTokenSecret, a.Cfg.DayOfToken)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &response.LoginResponse{
		AccessToken: token,
	}, nil
}

func (a *AccountsApp) UpdateUser(ctx context.Context, params *requests.UpdatePasswordBodyRequest) error {
	user := &models.User{}
	err := a.AccountsRepo.FindUser(ctx, user, param.FindParam{
		UserID: contextutil.GetUserID(ctx),
		Status: string(constant.Actived),
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.OldPassword))
	if err != nil {
		return tracerr.Wrap(err)
	}

	cryptPass, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), a.Cfg.IntBycrptPassword)
	if err != nil {
		return tracerr.Wrap(err)
	}

	// TODO: update password
	updateParam := param.UpdateParam{
		Password: string(cryptPass),
	}

	findParam := param.FindParam{
		UserID: user.Id,
	}

	err = a.AccountsRepo.UpdateUser(ctx, updateParam, findParam)
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func (a *AccountsApp) CreateNewAccount(ctx context.Context, params *requests.CreateAccountBodyRequest) error {
	cryptPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), a.Cfg.IntBycrptPassword)
	if err != nil {
		return tracerr.Wrap(err)
	}

	err = a.AccountsRepo.RunInTransaction(ctx, func(ctx context.Context) error {
		user := &models.User{
			Username:  params.Username,
			Password:  string(cryptPass),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// create super admin account with active status
		err = a.AccountsRepo.CreateUser(ctx, user)
		if err != nil {
			return tracerr.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}
