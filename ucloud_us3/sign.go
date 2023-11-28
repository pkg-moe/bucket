package ucloud_us3

import (
	"github.com/pkg/errors"
)

func (o *UCloudUS3) SignGetUrl(name string, expiredInSec int64) (string, error) {
	//auth := ufsdk.NewAuth(o.config.AK, o.config.SK)
	//
	//auth.AuthorizationPrivateURL("POST", o.config.Bucket, name, nil)

	return "", errors.New("not support")
}

func (o *UCloudUS3) SignPutUrl(name string, expiredInSec int64) (string, error) {
	return "", errors.New("not support")
}

func (o *UCloudUS3) SignPutUrlWithMD5(name string, expiredInSec int64, md5 string) (string, error) {
	//auth := ufsdk.NewAuth(o.config.AK, o.config.SK)
	//
	//h := make(http.Header)
	//h.Add("Content-Type", "application/octet-stream")
	//h.Add("Content-MD5", md5)
	//
	//signedURL := auth.Authorization("POST", o.config.Bucket, name, h)

	return "", errors.New("not support")
}
