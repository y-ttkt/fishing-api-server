package profile

import (
	"context"
	entity "github.com/yusuke-takatsu/fishing-api-server/domain/entity/profile"
	vo "github.com/yusuke-takatsu/fishing-api-server/domain/vo/profile"
	custumError "github.com/yusuke-takatsu/fishing-api-server/errors"
	"github.com/yusuke-takatsu/fishing-api-server/infra/repository/profile"
	"github.com/yusuke-takatsu/fishing-api-server/infra/repository/s3"
	dto "github.com/yusuke-takatsu/fishing-api-server/interface/dto/input/profile"
	"log"
)

type RegisterUseCase struct {
	profileRepo profile.Repository
	s3Repo      s3.Repository
}

func NewRegisterUseCase(profileRepo profile.Repository, s3Repo s3.Repository) *RegisterUseCase {
	return &RegisterUseCase{
		profileRepo: profileRepo,
		s3Repo:      s3Repo,
	}
}

func (s *RegisterUseCase) Execute(ctx context.Context, input dto.RegisterInputData) error {
	dob, err := vo.NewDateOfBirth(input.DateOfBirth)
	if err != nil {
		log.Printf("error creating dob %v", err)
		return custumError.Invalid.Wrap("無効な生年月日です。", err)
	}

	fsd, err := vo.NewFishingStartedDate(input.FishingStartedDate)
	if err != nil {
		log.Printf("error creating fishing started date %v", err)
		return custumError.Invalid.Wrap("無効な釣りを始めた日です。", err)
	}

	profileEntity, err := s.profileRepo.FindByUserID(ctx, input.UserID)
	if err != nil {
		log.Printf("error finding profile %v", err)
		return custumError.InternalErr.Wrap("プロフィールの取得に失敗しました。", err)
	}

	image, err := s.uploadFile(ctx, input, profileEntity)
	if err != nil {
		return err
	}

	if profileEntity != nil {
		profileEntity.Update(input.NickName, dob, fsd, image)
	}

	return s.profileRepo.UpdateOrCreate(ctx, profileEntity)
}

func (s *RegisterUseCase) uploadFile(ctx context.Context, input dto.RegisterInputData, entity *entity.Profile) (*vo.Image, error) {
	var imageVo *vo.Image

	if input.Image == nil {
		if entity.Image != nil {
			err := s.s3Repo.DeleteImage(ctx, entity.Image.Value())
			if err != nil {
				log.Printf("error deleting image %v", err)
				return nil, custumError.InternalErr.Wrap("画像の削除に失敗しました。", err)
			}
		}

		return nil, nil
	}

	if input.Image != nil {
		key, err := s.s3Repo.Upload(ctx, input.Image)
		if err != nil {
			log.Printf("error uploading image %v", err)
			return nil, custumError.InternalErr.Wrap("画像のアップロードに失敗しました。", err)
		}

		if entity.Image != nil {
			err = s.s3Repo.DeleteImage(ctx, entity.Image.Value())
			if err != nil {
				log.Printf("error deleting image %v", err)
				return nil, custumError.InternalErr.Wrap("画像の削除に失敗しました。", err)
			}
		}

		img := vo.NewImage(key)
		imageVo = &img
	}

	return imageVo, nil
}
