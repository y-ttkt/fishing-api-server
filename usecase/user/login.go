package user

import (
	"context"
	"database/sql"
	"errors"
	vo "github.com/yusuke-takatsu/fishing-api-server/domain/vo/user"
	custumError "github.com/yusuke-takatsu/fishing-api-server/errors"
	"github.com/yusuke-takatsu/fishing-api-server/infra/repository/user"
	dto "github.com/yusuke-takatsu/fishing-api-server/interface/dto/input/user"
)

type LoginUseCase struct {
	repo user.Repository
}

func NewLoginUseCase(repo user.Repository) *LoginUseCase {
	return &LoginUseCase{repo: repo}
}

func (s *LoginUseCase) Execute(ctx context.Context, input dto.LoginInputData) error {
	email, err := vo.NewEmail(input.Email)
	if err != nil {
		return custumError.Invalid.Wrap("無効なメールアドレスです。", err)
	}

	userEntity, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return custumError.NotFound.Wrap("ユーザーが見つかりませんでした。", err)
		}

		return custumError.InternalErr.Wrap("ユーザーの取得に失敗しました。", err)
	}

	err = vo.CompareHashAndPassword(userEntity.Password.Value(), input.Password)
	if err != nil {
		return custumError.Invalid.Wrap("パスワードが一致しません。", err)
	}

	return nil
}
