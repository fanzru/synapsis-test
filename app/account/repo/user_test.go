package repo_test

import (
	"context"
	"errors"
	"testing"

	"synapsis-test/app/account/domain/models"
	"synapsis-test/app/account/domain/param"

	"synapsis-test/app/account/repo"

	"synapsis-test/infrastructure/database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestFind(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %s", err)
	}
	defer mockDB.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error creating GORM DB instance: %s", err)
	}

	repo := &repo.AccountsRepo{
		MySQL: database.Connection{DB: gormDB},
	}

	tests := []struct {
		Name           string
		Param          param.FindParam
		WantErr        bool
		Expected       string
		TestCaseNumber int
	}{
		{
			Name: "Test case 1",
			Param: param.FindParam{
				Username: "testuser",
				UserID:   1,
			},
			WantErr:        false,
			Expected:       "testuser",
			TestCaseNumber: 1,
		},
		{
			Name: "Test case 2",
			Param: param.FindParam{
				Username: "testuser",
				Status:   "ACTIVED",
			},
			WantErr:        false,
			Expected:       "testuser",
			TestCaseNumber: 2,
		},
		{
			Name: "Test case 3",
			Param: param.FindParam{
				Username: "testuser",
				Status:   "ACTIVED",
			},
			WantErr:        true,
			Expected:       "testuser",
			TestCaseNumber: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if tt.TestCaseNumber == 1 {
				users := sqlmock.NewRows([]string{"username"}).AddRow(tt.Expected)
				mock.ExpectQuery("SELECT \\* FROM `users` WHERE users\\.id = \\? AND users\\.username = \\? AND `users`.`deleted_at` IS NULL").
					WithArgs(tt.Param.UserID, tt.Param.Username).
					WillReturnRows(users)
			}

			if tt.TestCaseNumber == 2 {
				users := sqlmock.NewRows([]string{"username"}).AddRow(tt.Expected)
				mock.ExpectQuery("SELECT \\* FROM `users` WHERE users\\.username = \\? AND users\\.status = \\? AND `users`.`deleted_at` IS NULL").
					WithArgs(tt.Param.Username, tt.Param.Status).
					WillReturnRows(users)
			}
			if tt.TestCaseNumber == 3 {

				mock.ExpectQuery("SELECT \\* FROM `users` WHERE users\\.username = \\? AND users\\.status = \\? AND `users`.`deleted_at` IS NULL").
					WithArgs(tt.Param.Username, tt.Param.Status).
					WillReturnError(errors.New("mock err"))
			}

			var foundUser models.User
			err := repo.FindUser(context.Background(), &foundUser, tt.Param)
			if tt.WantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected, foundUser.Username)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
