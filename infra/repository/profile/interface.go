package profile

import (
	"context"
	"github.com/yusuke-takatsu/fishing-api-server/domain/entity/profile"
)

type Repository interface {
	FindByUserID(ctx context.Context, userID string) (*profile.Profile, error)
	UpdateOrCreate(ctx context.Context, p *profile.Profile) error
}
