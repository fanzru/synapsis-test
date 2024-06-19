package repo

import (
	"context"
	"errors"
	"synapsis-test/app/account/domain/models"
	"synapsis-test/app/account/domain/param"
	"synapsis-test/app/account/domain/response"
	"synapsis-test/infrastructure/config"
	"synapsis-test/infrastructure/database"

	"gorm.io/gorm"
)

type Impl interface {
	// table : users
	DeleteUser(ctx context.Context, findParam param.FindParam) error
	FindFirstUser(ctx context.Context, user *models.User) (int64, error)
	FindUser(ctx context.Context, user *models.User, param param.FindParam) error
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, updateParam param.UpdateParam, findParam param.FindParam) error
	FindUserWithPagination(ctx context.Context, param param.FindParam) (response.Pagination, error)

	// for transactions
	RunInTransaction(ctx context.Context, f func(ctx context.Context) error) error
}
type AccountsRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(a *AccountsRepo) Impl {
	return a
}

type txnContextKey string

const (
	TableUsers               = "user"
	txnKey     txnContextKey = "transaction"
)

func (i *AccountsRepo) startTransaction(ctx context.Context) (context.Context, error) {
	tx := i.MySQL.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	txContext := context.WithValue(ctx, txnKey, tx)

	return txContext, nil
}

func (i *AccountsRepo) getTransaction(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txnKey).(*gorm.DB); ok {
		return tx
	}
	return nil
}

func (i *AccountsRepo) getMysqlConnection(ctx context.Context) *gorm.DB {
	db := i.getTransaction(ctx)
	if db == nil {
		db = i.MySQL.DB
	}
	return db
}

func (i *AccountsRepo) RunInTransaction(ctx context.Context, f func(ctx context.Context) error) error {
	txContext, err := i.startTransaction(ctx)
	if err != nil {
		return err
	}

	// Dapatkan transaksi dari context
	tx := i.getTransaction(txContext)
	if tx == nil {
		return errors.New("failed to get transaction from context")
	}

	// Jalankan fungsi f dengan context yang telah memiliki transaksi
	err = f(txContext)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaksi jika tidak ada error
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
