package cos

import (
	"context"
	"sort"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"

	bucketdao "pkg.moe/pkg/bucket/dao"
)

func (o *COS) ListObjectsV2(path string) ([]bucketdao.ObjectInfo, error) {
	res := []bucketdao.ObjectInfo{}

	continueToken := ""
	for {
		list, _, err := o.client.Bucket.Get(context.TODO(), &cos.BucketGetOptions{
			Prefix:    path,
			Delimiter: "/",
			MaxKeys:   1000,
			Marker:    continueToken,
		})
		if err != nil {
			return nil, err
		}

		for _, item := range list.Contents {
			t, err := time.Parse(time.RFC3339, item.LastModified)
			if err != nil {
				return nil, err
			}

			res = append(res, bucketdao.ObjectInfo{
				Key:          item.Key,
				LastModified: t.Unix(),
				Size:         item.Size,
			})
		}

		if list.IsTruncated {
			continueToken = list.NextMarker
		} else {
			break
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].LastModified > res[j].LastModified
	})

	return res, nil
}
