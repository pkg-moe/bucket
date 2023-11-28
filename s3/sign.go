package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (o *S3) SignGetUrl(name string, expiredInSec int64) (string, error) {
	signedURL, err := o.signer.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(o.config.Bucket),
		Key:    aws.String(name),
	}, func(options *s3.PresignOptions) {
		options.Expires = time.Duration(expiredInSec) * time.Second
	})

	return signedURL.URL, err
}

func (o *S3) SignPutUrl(name string, expiredInSec int64) (string, error) {
	signedURL, err := o.signer.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(o.config.Bucket),
		Key:    aws.String(name),
	}, func(options *s3.PresignOptions) {
		options.Expires = time.Duration(expiredInSec) * time.Second
	})

	return signedURL.URL, err
}

func (o *S3) SignPutUrlWithMD5(name string, expiredInSec int64, md5 string) (string, error) {
	signedURL, err := o.signer.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:     aws.String(o.config.Bucket),
		Key:        aws.String(name),
		ContentMD5: aws.String(md5),
	}, func(options *s3.PresignOptions) {
		options.Expires = time.Duration(expiredInSec) * time.Second
	})

	return signedURL.URL, err
}
