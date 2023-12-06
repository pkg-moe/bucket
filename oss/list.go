package oss

import (
	"fmt"
	"sort"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	bucketdao "pkg.moe/pkg/bucket/dao"
)

func (o *OSS) ListObjectsV2(path string) ([]bucketdao.ObjectInfo, error) {

	res := make([]bucketdao.ObjectInfo, 0)
	continueToken := ""
	for {
		lsRes, err := o.bucket.ListObjectsV2(oss.Prefix(path), oss.ContinuationToken(continueToken))
		if err != nil {
			return nil, err
		}

		// 打印列举文件，默认情况下，一次返回100条记录。
		for _, object := range lsRes.Objects {
			fmt.Println(object.Key, object.Type, object.Size, object.ETag, object.LastModified, object.StorageClass)
			res = append(res, bucketdao.ObjectInfo{
				Key:          object.Key,
				LastModified: object.LastModified.Unix(),
				Size:         object.Size,
			})
		}

		if lsRes.IsTruncated {
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].LastModified > res[j].LastModified
	})

	return res, nil
}
