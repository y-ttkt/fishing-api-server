package user

import (
	"github.com/yusuke-takatsu/fishing-api-server/domain/vo/user"
	enum "github.com/yusuke-takatsu/fishing-api-server/enum/user"
	"github.com/yusuke-takatsu/fishing-api-server/util"
	"time"
)

type User struct {
	ID        string
	Email     user.Email
	Password  user.Password
	Status    enum.Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newUser(
	id string,
	email user.Email,
	password user.Password,
	status enum.Status,
) *User {
	now := time.Now()

	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		Status:    status,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func CreateFromDB(
	id string,
	email user.Email,
	password user.Password,
	status enum.Status,
) *User {
	return newUser(
		id,
		email,
		password,
		status,
	)
}

func CreateUser(
	email user.Email,
	password user.Password,
) *User {
	return newUser(
		util.GenerateIdentifier().Identifier,
		email,
		password,
		enum.ProvisionalMember,
	)
}
