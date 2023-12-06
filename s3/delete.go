package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (o *S3) DeleteObjects(keys []string) (err error) {
	l := []types.ObjectIdentifier{}
	for _, key := range keys {
		l = append(l, types.ObjectIdentifier{
			Key: aws.String(key),
		})
	}

	if len(l) <= 1000 {
		_, err = o.client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
			Bucket: aws.String(o.config.Bucket),
			Delete: &types.Delete{
				Objects: l,
				Quiet:   aws.Bool(true),
			},
		})
	} else {
		from := 0
		hasMore := true
		for hasMore {
			end := from + 1000
			if end > len(l) {
				end = len(l)
				hasMore = false
			}
			ss := l[from:end]
			_, err = o.client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
				Bucket: aws.String(o.config.Bucket),
				Delete: &types.Delete{
					Objects: ss,
					Quiet:   aws.Bool(true),
				},
			})
			if err != nil {
				return err
			}
			from = end
		}

	}

	return
}
