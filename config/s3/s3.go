package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"os"
	"strconv"
)

var s3Client *s3.Client

func InitS3Client() *s3.Client {
	creds := credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(os.Getenv("AWS_DEFAULT_REGION")),
		config.WithBaseEndpoint(os.Getenv("AWS_ENDPOINT")),
		config.WithCredentialsProvider(creds),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	usePathStyle, err := strconv.ParseBool(os.Getenv("AWS_USE_PATH_STYLE_ENDPOINT"))
	if err != nil {
		log.Fatalf("unable to parse AWS_USE_PATH_STYLE_ENDPOINT, %v", err)
	}

	s3Client = s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = usePathStyle
	})

	return s3Client
}
