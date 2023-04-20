package cos

import (
	"context"
	"net/http"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func (c *COS) SignGetUrl(name string, expiredInSec int64) (string, error) {
	signedURL, err := c.client.Object.GetPresignedURL(context.Background(), http.MethodGet, name, c.config.AK, c.config.SK, time.Duration(expiredInSec)*time.Second, nil)
	if err != nil {
		return "", err
	}
	return signedURL.String(), err
}

func (c *COS) SignPutUrl(name string, expiredInSec int64) (string, error) {
	signedURL, err := c.client.Object.GetPresignedURL(context.Background(), http.MethodPut, name, c.config.AK, c.config.SK, time.Duration(expiredInSec)*time.Second, nil)
	if err != nil {
		return "", err
	}
	return signedURL.String(), err
}

func (c *COS) SignPutUrlWithMD5(name string, expiredInSec int64, md5 string) (string, error) {
	signedURL, err := c.client.Object.GetPresignedURL(context.Background(), http.MethodPut, name, c.config.AK, c.config.SK, time.Duration(expiredInSec)*time.Second, &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentMD5: md5,
		},
	})
	if err != nil {
		return "", err
	}
	return signedURL.String(), err
}
