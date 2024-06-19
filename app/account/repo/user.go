package repo

import (
	"context"
	"errors"
	"synapsis-test/app/account/domain/constant"
	"synapsis-test/app/account/domain/models"
	"synapsis-test/app/account/domain/param"
	"synapsis-test/app/account/domain/response"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

func (i *AccountsRepo) UpdateUser(ctx context.Context, updateParam param.UpdateParam, findParam param.FindParam) error {
	db := i.MySQL.DB.Table(TableUsers)

	db = i.selectBuildersUser(findParam, db)

	user := i.updatesBuildersUser(updateParam)

	result := db.Updates(user)
	if result.Error != nil {
		return tracerr.Wrap(result.Error)
	}

	return nil
}

func (i *AccountsRepo) CreateUser(ctx context.Context, user *models.User) error {
	db := i.getMysqlConnection(ctx)
	result := db.Table(TableUsers).Create(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrDuplicatedKey {
			return tracerr.Wrap(constant.ErrCantCreateUser)
		}
		return tracerr.Wrap(result.Error)
	}

	return nil
}

func (i *AccountsRepo) FindUser(ctx context.Context, user *models.User, param param.FindParam) error {
	db := i.MySQL.DB.Table(TableUsers)

	db = i.selectBuildersUser(param, db)

	result := db.Find(&user)
	if result.Error != nil {
		return tracerr.Wrap(result.Error)
	}

	return nil
}

func (i *AccountsRepo) FindFirstUser(ctx context.Context, user *models.User) (int64, error) {
	db := i.MySQL.DB.Table(TableUsers)

	result := db.First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return -1, tracerr.Wrap(result.Error)
	}

	return db.RowsAffected, nil
}

func (i *AccountsRepo) DeleteUser(ctx context.Context, findParam param.FindParam) error {
	user := &models.User{
		Id: findParam.UserID,
	}
	db := i.MySQL.DB.Table(TableUsers)

	db = i.selectBuildersUser(findParam, db)

	result := db.First(user)
	if result.Error != nil {
		return tracerr.Wrap(result.Error)
	}

	result = i.MySQL.DB.Table(TableUsers).Delete(user)
	if result.Error != nil {
		return tracerr.Wrap(result.Error)
	}

	return nil
}

func (i *AccountsRepo) FindUserWithPagination(ctx context.Context, param param.FindParam) (response.Pagination, error) {
	resp := response.Pagination{}
	var totalRows int64

	db := i.MySQL.DB.Table(TableUsers)

	db = i.selectBuildersUser(param, db)

	if err := db.Count(&totalRows).Error; err != nil {
		return resp, err
	}

	pageSize := 10 // Ukuran halaman
	offset := (param.PageNumber - 1) * pageSize

	users := []*models.UserWithoutPassword{}
	if err := db.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return resp, err
	}

	totalPages := (totalRows + int64(pageSize) - 1) / int64(pageSize)

	var prevPage, nextPage int
	if param.PageNumber > 1 {
		prevPage = param.PageNumber - 1
	}
	if offset+pageSize < int(totalRows) {
		nextPage = param.PageNumber + 1
	}

	resp = response.Pagination{
		Page:  param.PageNumber,
		Limit: int64(pageSize),
		Prev:  int64(prevPage),
		Next:  int64(nextPage),
		Start: 1,
		End:   totalPages,
		Data:  users,
	}

	return resp, nil
}

func (i *AccountsRepo) selectBuildersUser(param param.FindParam, db *gorm.DB) *gorm.DB {

	if param.UserID > 0 {
		db = db.Where("user.id = ?", param.UserID)
	}

	if param.Username != "" {
		db = db.Where("user.username = ?", param.Username)
	}

	return db
}

func (i *AccountsRepo) updatesBuildersUser(param param.UpdateParam) map[string]interface{} {
	user := make(map[string]interface{})

	if param.Password != "" {
		user["password"] = param.Password
	}

	return user
}
