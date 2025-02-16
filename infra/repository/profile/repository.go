package profile

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yusuke-takatsu/fishing-api-server/domain/entity/profile"
	vo "github.com/yusuke-takatsu/fishing-api-server/domain/vo/profile"
	"time"
)

type profileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) Repository {
	return &profileRepository{db: db}
}

func (r *profileRepository) FindByUserID(ctx context.Context, userID string) (*profile.Profile, error) {
	query := `
		SELECT nick_name, date_of_birth, fishing_started_date, image
		FROM profiles
		WHERE user_id = ?
	`

	row := r.db.QueryRowContext(ctx, query, userID)
	var nickName, dob, fsd string
	var image sql.NullString
	err := row.Scan(&nickName, &dob, &fsd, &image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	dobVo, err := vo.NewDateOfBirth(dob)
	if err != nil {
		return nil, err
	}

	fsdVo, err := vo.NewFishingStartedDate(fsd)
	if err != nil {
		return nil, err
	}

	var imageVo *vo.Image
	if image.Valid {
		img := vo.NewImage(image.String)
		imageVo = &img
	}

	return profile.NewProfile(
		userID,
		nickName,
		dobVo,
		fsdVo,
		imageVo,
	), nil
}

func (r *profileRepository) UpdateOrCreate(ctx context.Context, p *profile.Profile) error {
	updateQuery := `
		UPDATE profiles
		SET nick_name = ?, date_of_birth = ?, fishing_started_date = ?, image = ?, updated_at = ?
		where user_id = ?
	`

	res, err := r.db.ExecContext(ctx, updateQuery,
		p.NickName,
		p.DateOfBirth,
		p.FishingStartedDate,
		p.Image,
		time.Now(),
		p.UserID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// 存在しなければ INSERT
	if rowsAffected == 0 {
		queryInsert := `
			INSERT INTO profiles 
			(user_id, nick_name, date_of_birth, fishing_started_date, image, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`
		_, err = r.db.ExecContext(ctx, queryInsert,
			p.UserID,
			p.NickName,
			p.DateOfBirth.Value().Format("2006-01-02"),
			p.FishingStartedDate.Value().Format("2006-01-02"),
			p.Image,
			time.Now(),
			time.Now(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
