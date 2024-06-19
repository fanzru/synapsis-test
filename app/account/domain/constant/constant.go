package constant

import (
	"errors"
)

var (
	// errors code
	ErrCantCreateSuperAdmin = errors.New("database user already have value, doesn't create super admin")
	ErrCantCreateUser       = errors.New("username already exist")
	ErrWrongPassword        = errors.New("wrong password")
)

type AccountType string
type AccountStatus string

const (
	SuperAdmin AccountType = "SUPER_ADMIN"
	User       AccountType = "USER"
)

const (
	Pending AccountStatus = "PENDING"
	Actived AccountStatus = "ACTIVED"
)
