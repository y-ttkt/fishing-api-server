package user

import (
	"context"
	entity "github.com/yusuke-takatsu/fishing-api-server/domain/entity/user"
	"github.com/yusuke-takatsu/fishing-api-server/domain/vo/user"
)

type Repository interface {
	FindByEmail(ctx context.Context, email user.Email) (*entity.User, error)
}
