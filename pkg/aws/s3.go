package aws

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Bucket struct {
	ctx    context.Context
	sess   *session.Session
	bucket string
}

func NewS3(ctx context.Context, bucket string, sess *session.Session) *S3Bucket {
	return &S3Bucket{
		ctx:    ctx,
		sess:   sess,
		bucket: bucket,
	}
}

func (s *S3Bucket) UploadObject(key string, file multipart.File) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(s.sess)

	output, err := uploader.UploadWithContext(s.ctx, &s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(key),
		Body:   file,
	})

	return output, err
}
