package s3

import (
	"context"
	"mime/multipart"
)

type Repository interface {
	Upload(ctx context.Context, fileHeader *multipart.FileHeader) (string, error)
	DeleteImage(ctx context.Context, key string) error
}
