package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	s3manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/yusuke-takatsu/fishing-api-server/domain/vo/profile"
	"mime/multipart"
	"os"
	"path/filepath"
)

type S3Repository struct {
	s3Client *s3.Client
}

func NewS3Repository(s3Client *s3.Client) Repository {
	return &S3Repository{s3Client: s3Client}
}

func (s *S3Repository) Upload(ctx context.Context, fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	hash := profile.MakeHashName(ext)
	key := fmt.Sprintf("images/%s", hash.Value())

	uploader := s3manager.NewUploader(s.s3Client)
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_PUBLIC_BUCKET")),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	return key, nil
}

func (s *S3Repository) DeleteImage(ctx context.Context, key string) error {
	_, err := s.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_PUBLIC_BUCKET")),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}

	return nil
}
