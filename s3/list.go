package s3

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	bucketdao "pkg.moe/pkg/bucket/dao"
)

func (o *S3) ListObjectsV2(path string) ([]bucketdao.ObjectInfo, error) {
	res := []bucketdao.ObjectInfo{}

	list, err := o.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:    aws.String(o.config.Bucket),
		Prefix:    aws.String(path),
		Delimiter: aws.String("/"),
		MaxKeys:   aws.Int32(1000),
	})
	if err != nil {
		return nil, err
	}

	for _, item := range list.Contents {
		res = append(res, bucketdao.ObjectInfo{
			Key:          *item.Key,
			LastModified: item.LastModified.Unix(),
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].LastModified > res[j].LastModified
	})

	return res, nil
}
