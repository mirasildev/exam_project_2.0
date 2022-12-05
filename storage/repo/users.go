package repo

import "time"

const (
	UserTypeSuperadmin = "superadmin"
	UserTypeUser       = "user"
	UserTypePartner    = "partner"
)

type User struct {
	ID          int64
	FirstName   string
	LastName    string
	PhoneNumber *string
	Email       string
	Password    string
	Type        string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

type GetAllUsersParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllUsersResult struct {
	Users []*User
	Count int32
}

type UpdatePassword struct {
	UserID   int64
	Password string
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(id int64) (*User, error)
	GetAll(params *GetAllUsersParams) (*GetAllUsersResult, error)
	Update(u *User) (*User, error)
	UpdatePassword(req *UpdatePassword) error
	Delete(id int64) error
	GetByEmail(email string) (*User, error)
}
