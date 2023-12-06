package oss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func (o *OSS) DeleteObjects(keys []string) (err error) {
	if len(keys) <= 1000 {
		_, err = o.bucket.DeleteObjects(keys, oss.DeleteObjectsQuiet(true))
	} else {

		from := 0
		hasMore := true
		for hasMore {
			end := from + 1000
			if end > len(keys) {
				end = len(keys)
				hasMore = false
			}
			ss := keys[from:end]
			_, err = o.bucket.DeleteObjects(ss, oss.DeleteObjectsQuiet(true))
			if err != nil {
				return err
			}
			from = end
		}

	}

	return
}
