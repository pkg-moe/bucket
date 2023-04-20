package cos

import (
	"net/http"
	"net/url"

	"pkg.moe/pkg/bucket/dao"
)

type COS struct {
	config *bucketdao.Config

	client *cos.Client
}

func (c *COS) Init(config *bucketdao.Config) (err error) {
	c.config = config

	bucketUrl := "https://" + config.Bucket + ".cos." + config.Endpoint + ".myqcloud.com"

	u, err := url.Parse(bucketUrl)
	if err != nil {
		return err
	}

	b := &cos.BaseURL{BucketURL: u}

	c.client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.AK,
			SecretKey: config.SK,
		},
	})

	return
}
