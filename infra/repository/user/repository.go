package user

import (
	"context"
	"database/sql"
	entity "github.com/yusuke-takatsu/fishing-api-server/domain/entity/user"
	"github.com/yusuke-takatsu/fishing-api-server/domain/vo/user"
	enum "github.com/yusuke-takatsu/fishing-api-server/enum/user"
)

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(ctx context.Context, email user.Email) (*entity.User, error) {
	query := "SELECT id, email, password, status FROM users WHERE email = ? LIMIT 1"
	row := r.db.QueryRowContext(ctx, query, email.Value())

	var id string
	var emailStr string
	var password string
	var status int

	if err := row.Scan(&id, &emailStr, &password, &status); err != nil {
		return nil, err
	}

	emailVo, err := user.NewEmail(emailStr)
	if err != nil {
		return nil, err
	}

	return entity.CreateFromDB(id, emailVo, user.CreateFromDB(password), enum.Status(status)), nil
}
