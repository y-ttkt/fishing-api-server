package profile

import "mime/multipart"

type RegisterInputData struct {
	UserID             string
	NickName           string
	DateOfBirth        string
	FishingStartedDate string
	Image              *multipart.FileHeader
}
