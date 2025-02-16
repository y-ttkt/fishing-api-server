package profile

import (
	"github.com/yusuke-takatsu/fishing-api-server/domain/vo/profile"
)

type Profile struct {
	UserID             string
	NickName           string
	DateOfBirth        profile.DateOfBirth
	FishingStartedDate profile.FishingStartedDate
	Image              *profile.Image
}

func NewProfile(
	userID, nickName string,
	dob profile.DateOfBirth,
	fsd profile.FishingStartedDate,
	img *profile.Image,
) *Profile {
	return &Profile{
		UserID:             userID,
		NickName:           nickName,
		DateOfBirth:        dob,
		FishingStartedDate: fsd,
		Image:              img,
	}
}

func (p *Profile) Update(
	nickName string,
	dob profile.DateOfBirth,
	fsd profile.FishingStartedDate,
	img *profile.Image,
) {
	p.NickName = nickName
	p.DateOfBirth = dob
	p.FishingStartedDate = fsd
	p.Image = img
}
