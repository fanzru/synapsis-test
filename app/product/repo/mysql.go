package repo

import (
	"context"
	"errors"
	"synapsis-test/app/product/domain/models"
	"synapsis-test/app/product/domain/params"
	"synapsis-test/infrastructure/config"
	"synapsis-test/infrastructure/database"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type Impl interface {
	FindProducts(ctx context.Context, params params.FindParams) (*[]models.Product, error)
	FindCategory(ctx context.Context, params params.FindParams) (*[]models.Category, error)
}
type ProductRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(a *ProductRepo) Impl {
	return a
}

type txnContextKey string

const (
	TableProduct                  = "product"
	TableCategories               = "category"
	txnKey          txnContextKey = "transaction"
)

func (i *ProductRepo) startTransaction(ctx context.Context) (context.Context, error) {
	tx := i.MySQL.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	txContext := context.WithValue(ctx, txnKey, tx)

	return txContext, nil
}

func (i *ProductRepo) getTransaction(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txnKey).(*gorm.DB); ok {
		return tx
	}
	return nil
}

// func (i *ProductRepo) getMysqlConnection(ctx context.Context) *gorm.DB {
// 	db := i.getTransaction(ctx)
// 	if db == nil {
// 		db = i.MySQL.DB
// 	}
// 	return db
// }

func (i *ProductRepo) RunInTransaction(ctx context.Context, f func(ctx context.Context) error) error {
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

func (i *ProductRepo) FindProducts(ctx context.Context, params params.FindParams) (*[]models.Product, error) {
	var data *[]models.Product
	db := i.MySQL.DB.Preload("Category").Table(TableProduct)

	if params.CategoryID != 0 {
		db = db.Joins("Category").Where("Category.id = ?", params.CategoryID)
	}

	result := db.Find(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tracerr.Wrap(result.Error)
	}

	return data, nil
}

func (i *ProductRepo) FindCategory(ctx context.Context, params params.FindParams) (*[]models.Category, error) {
	var data *[]models.Category
	db := i.MySQL.DB.Preload("Products").Table(TableCategories)

	if params.CategoryID != 0 {
		db = db.Where("id = ?", params.CategoryID)
	}

	result := db.Find(&data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tracerr.Wrap(result.Error)
	}

	return data, nil
}
