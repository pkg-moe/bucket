package ucloud_us3

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
)

func (o *UCloudUS3) SignGetUrl(name string, expiredInSec int64) (string, error) {
	req, err := ufsdk.NewFileRequest(o.clientConfig, nil)
	if err != nil {
		return "", err
	}

	signUrl := req.GetPrivateURL(name, time.Duration(expiredInSec)*time.Minute)

	return signUrl, nil
}

func (o *UCloudUS3) SignPutUrl(name string, expiredInSec int64) (string, error) {
	auth := ufsdk.NewAuth(o.config.AK, o.config.SK)

	t := time.Now()
	t = t.Add(time.Duration(expiredInSec) * time.Second)
	expires := strconv.FormatInt(t.Unix(), 10)

	sign, publicKey := auth.AuthorizationPrivateURL("PUT", o.config.Bucket, name, expires, nil)

	req, err := ufsdk.NewFileRequest(o.clientConfig, nil)
	if err != nil {
		return "", err
	}

	query := url.Values{}
	query.Add("UCloudPublicKey", publicKey)
	query.Add("Signature", sign)
	query.Add("Expires", expires)
	reqURL := req.GetPublicURL(name)
	return reqURL + "?" + query.Encode(), nil
}

func (o *UCloudUS3) SignPutUrlWithMD5(name string, expiredInSec int64, md5 string) (string, error) {
	auth := ufsdk.NewAuth(o.config.AK, o.config.SK)

	t := time.Now()
	t = t.Add(time.Duration(expiredInSec) * time.Second)
	expires := strconv.FormatInt(t.Unix(), 10)

	h := http.Header{}
	h.Set("Content-MD5", md5)

	sign, publicKey := auth.AuthorizationPrivateURL("PUT", o.config.Bucket, name, expires, h)

	req, err := ufsdk.NewFileRequest(o.clientConfig, nil)
	if err != nil {
		return "", err
	}

	query := url.Values{}
	query.Add("UCloudPublicKey", publicKey)
	query.Add("Signature", sign)
	query.Add("Expires", expires)
	reqURL := req.GetPublicURL(name)

	return reqURL + "?" + query.Encode(), nil
}
