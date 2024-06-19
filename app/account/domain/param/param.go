package param

type FindParam struct {
	UserID   int64
	Username string
	Password string
	Status   string
	RoleName string

	// for pagination
	PageNumber int
}

type UpdateParam struct {
	Username string
	Password string
	Status   string
}

type DeleteParam struct {
	Username string
	Password string
	Status   string
}

type CreateAccountParam struct {
	Username string
	Password string
}
