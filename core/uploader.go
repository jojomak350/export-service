package core

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type Uploader struct {
	session *session.Session
	client  *s3manager.Uploader
}

var UploaderClient Uploader

func NewUploader() {
	s := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""),
	}))

	UploaderClient = Uploader{
		session: s,
		client:  s3manager.NewUploader(s),
	}
}

func (u *Uploader) cleanUp(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}

	if err := os.Remove(file.Name()); err != nil {
		panic(err)
	}
}

func (u *Uploader) Upload(file *os.File, path string) string {
	out, err := u.client.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(path),
		Body:   file,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		panic(err)
	}

	u.cleanUp(file)

	return out.Location
}
