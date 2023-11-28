package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"pkg.moe/pkg/bucket/dao"
)

type S3 struct {
	config *bucketdao.Config

	client *s3.Client
	signer *s3.PresignClient
}

func (o *S3) Init(c *bucketdao.Config) (err error) {
	o.config = c

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   c.PartitionID,
			URL:           c.Endpoint,
			SigningRegion: c.Region,
		}, nil
	})

	creds := credentials.NewStaticCredentialsProvider(c.AK, c.SK, "")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		return err
	}

	o.client = s3.NewFromConfig(cfg)

	o.signer = s3.NewPresignClient(o.client)

	return nil
}
