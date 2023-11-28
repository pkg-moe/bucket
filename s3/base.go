package s3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (o *S3) Get(name string) ([]byte, error) {
	m := manager.NewDownloader(o.client)

	buf := manager.NewWriteAtBuffer([]byte{})
	_, err := m.Download(context.TODO(), buf, &s3.GetObjectInput{
		Bucket: aws.String(o.config.Bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), err
}

func (o *S3) Put(name string, value []byte) error {
	m := manager.NewUploader(o.client)

	_, err := m.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(o.config.Bucket),
		Key:    aws.String(name),
		Body:   bytes.NewReader(value),
	})
	if err != nil {
		return err
	}

	return nil
}
