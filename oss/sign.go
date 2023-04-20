package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func (o *OSS) SignGetUrl(name string, expiredInSec int64) (string, error) {
	signedURL, err := o.bucket.SignURL(name, oss.HTTPGet, expiredInSec)
	return signedURL, err
}

func (o *OSS) SignPutUrl(name string, expiredInSec int64) (string, error) {
	signedURL, err := o.bucket.SignURL(name, oss.HTTPPut, expiredInSec)
	return signedURL, err
}

func (o *OSS) SignPutUrlWithMD5(name string, expiredInSec int64, md5 string) (string, error) {
	signedURL, err := o.bucket.SignURL(name, oss.HTTPPut, expiredInSec, oss.ContentMD5(md5))
	return signedURL, err
}
