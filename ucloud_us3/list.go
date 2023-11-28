package ucloud_us3

import (
	"sort"

	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"

	bucketdao "pkg.moe/pkg/bucket/dao"
)

func (o *UCloudUS3) ListObjectsV2(path string) ([]bucketdao.ObjectInfo, error) {

	res := make([]bucketdao.ObjectInfo, 0)

	req, err := ufsdk.NewFileRequest(o.clientConfig, nil)
	if err != nil {
		return nil, err
	}

	list, err := req.ListObjects(path, "", "/", 1000)
	if err != nil {
		return nil, err
	}

	for _, item := range list.Contents {
		res = append(res, bucketdao.ObjectInfo{
			Key:          item.Key,
			LastModified: int64(item.LastModified),
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].LastModified > res[j].LastModified
	})

	return res, nil
}
