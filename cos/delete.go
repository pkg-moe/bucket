package cos

import (
	"context"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func (o *COS) DeleteObjects(keys []string) (err error) {
	l := []cos.Object{}
	for _, key := range keys {
		l = append(l, cos.Object{
			Key: key,
		})
	}

	if len(l) <= 1000 {
		_, _, err = o.client.Object.DeleteMulti(context.TODO(), &cos.ObjectDeleteMultiOptions{
			Quiet:   true,
			Objects: l,
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
			_, _, err = o.client.Object.DeleteMulti(context.TODO(), &cos.ObjectDeleteMultiOptions{
				Quiet:   true,
				Objects: ss,
			})
			if err != nil {
				return err
			}
			from = end
		}

	}

	return
}
