package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"pkg.moe/pkg/bucket/dao"
)

type OSS struct {
	config *bucketdao.Config

	client *oss.Client
	bucket *oss.Bucket
}

func (o *OSS) Init(config *bucketdao.Config) (err error) {
	o.config = config

	o.client, err = oss.New(config.Endpoint, config.AK, config.SK)
	if err != nil {
		return
	}
	o.bucket, err = o.client.Bucket(config.Bucket)
	return
}
